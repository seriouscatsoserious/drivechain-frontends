import 'dart:ui';

/// References the color scheme
abstract class SailColorScheme {
  // both theme colors
  static const blue = Color(0xff16A1F6);
  static const pinkLight = Color(0xffFEC1D5);
  static const red = Color(0xffEB5757);
  static const black = Color(0xff000000);
  static const blackLighter = Color(0xff0F172A);
  static const blackLightest = Color(0xff334155);
  static const greenLight = Color(0xffB5E5AE);
  static const green = Color(0xff4CB782);
  static const orange = Color(0xffEDB142);
  static const orangeLight = Color(0xffFFDE99);
  static const purple = Color(0xffBB87FC);
  static const white = Color(0xffFFFFFF);
  static const whiteDark = Color(0xffE7E7E7);
  static const greyDark = Color.fromRGBO(115, 115, 115, 1);
  static const greyMiddle = Color(0xffDADADA);
  static const greyLight = Color(0xff737373);
  static const blackLight = Color.fromARGB(255, 18, 18, 18);
  static const blackTertiary = Color(0xff21232E);
  static const superLightGreen = Color(0xffEEF8ED);
  static const lightGreen = Color(0xffEEF8ED);
  static const mediumGreen = Color(0xff5CBE45);
  static const mediumDarkGreen = Color(0xff59B23A);
  static const formFieldGrey = Color(0xffE9EBF0);
  static const whiteLight = Color(0xffEEEFFC);

  // dark theme colors
  static const darkBackground = Color(0xff181921);
  static const darkActionHeader = Color(0xff21232E);
  static const darkDivider = Color(0xff212234);
  static const darkChip = Color.fromRGBO(124, 124, 164, 0.1); // #7C7CA4
  static const darkTextSecondary = Color(0xffD2D3E0);
  static const darkTextTertiary = Color(0xff858699);
  static const darkInputBorder = Color(0xff393A4B);
  // modal colors
  static const darkActionModalBackground = Color.fromRGBO(29, 30, 43, 1);
  static const darkTextHint = Color(0xff4D4F69);
  // light theme colors
}
