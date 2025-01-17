import 'dart:io';

import 'package:get_it/get_it.dart';
import 'package:logger/logger.dart';
import 'package:sail_ui/config/binaries.dart';
import 'package:sail_ui/sail_ui.dart';

class Config {
  final String path;
  final String host;
  final int port;
  final String username;
  final String password;

  const Config({
    required this.path,
    required this.host,
    required this.port,
    required this.username,
    required this.password,
  });
}

// Order of precedence:
//
// 1. Conf file
// 2. Inspect cookie
// 3. Defaults
//
Future<NodeConnectionSettings> readRPCConfig(
  String datadir,
  String confFile,
  Binary chain,
  String network, {
  // if set, will force this network, irregardless of runtime argument
  bool useCookieAuth = true,
}) async {
  final log = GetIt.I.get<Logger>();

  final conf = File(filePath([datadir, confFile]));
  // Mainnet == empty string, special datadirs only apply to non-mainnet
  final networkDir = filePath([datadir, network == 'mainnet' ? '' : network]);

  final cookie = File(filePath([networkDir, '.cookie']));

  String? username;
  String? password;
  String? host;
  int? port;

  if (!await conf.exists() && !await cookie.exists()) {
    log.d('missing both conf ($conf) and cookie ($cookie), returning default settings');
    return NodeConnectionSettings(
      conf.path,
      'localhost',
      chain.port,
      'user',
      'password',
      network == 'regtest',
    );
  }

  if (useCookieAuth && await cookie.exists()) {
    final data = await cookie.readAsString();
    final parts = data.split(':');
    if (parts.length != 2) {
      throw 'unexpected cookie file: $data';
    }
    [username, password] = parts;

    log.t('rpc: read password from cookie file at $cookie');
  }

  if (await conf.exists()) {
    log.d('rpc: reading conf file at $conf');
    final confContent = await conf.readAsString();
    final lines = confContent.split('\n').map((e) => e.trim()).toList();

    // Create settings with basic values
    final settings = NodeConnectionSettings(
      conf.path,
      'localhost',
      chain.port,
      'user',
      'password',
      network == 'regtest',
    );

    // Read all config values, overwrite
    // default 'host' 'port' 'user' 'password' if set
    settings.readConfigFromFile(lines);

    return settings;
  }

  if (password == null || username == null) {
    throw 'could not find username or password in conf file, and cookie was not present';
  }

  host ??= 'localhost';
  port ??= chain.port;

// Make sure to not include password here
  log.i('resolved conf: $username@$host:$port on $network');
  return NodeConnectionSettings(
    conf.path,
    host,
    port,
    username,
    password,
    network == 'regtest',
  );
}

List<String> bitcoinCoreBinaryArgs(NodeConnectionSettings conf) {
  return [
    conf.isLocalNetwork ? '-regtest' : '',
    conf.username.isNotEmpty ? '-rpcuser=${conf.username}' : '',
    conf.password.isNotEmpty ? '-rpcpassword=${conf.password}' : '',
    conf.port != 0 ? '-rpcport=${conf.port}' : '',
    ...conf.getConfigArgs(), // Add all additional config values
  ]
      // important: empty strings trip up the binary
      .where((arg) => arg.isNotEmpty)
      .toList();
}

// checks if the loaded bitcoin core config contains a specific
// key, e.g:
// # testchain.conf
// regtest=1
//
// confKeyExists(args, 'regtest') => true
bool _confKeyExists(List<String> args, String key) {
  return args.any((arg) => arg.replaceAll('-', '').split('=').first == key);
}

void addEntryIfNotSet(List<String> args, String key, String value) {
  Logger log = GetIt.I.get<Logger>();

  if (_confKeyExists(args, key)) {
    return;
  }

  // args are expected on the form -paramsdir=/home/.zcash
  final newEntry = '-$key=$value';
  log.i('$key not present in conf, adding $newEntry');
  args.add(newEntry);
}
