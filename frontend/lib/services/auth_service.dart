import 'package:flutter/foundation.dart';

class AuthService with ChangeNotifier {
  bool _isAuthenticated = false;

  bool get isAuthenticated => _isAuthenticated;

  Future<void> login(String email, String password) async {
    // Заглушка для аутентификации
    await Future.delayed(Duration(seconds: 1));
    _isAuthenticated = true;
    notifyListeners();
  }

  Future<void> logout() async {
    _isAuthenticated = false;
    notifyListeners();
  }
}