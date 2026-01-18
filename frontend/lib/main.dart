import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'screens/login_screen.dart';
import 'screens/home_screen.dart';
import 'services/auth_service.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (context) => AuthService(),
      child: MaterialApp(
        title: 'UnityComm',
        theme: ThemeData(
          primarySwatch: Colors.blue,
        ),
        home: Consumer<AuthService>(
          builder: (context, auth, child) {
            return auth.isAuthenticated ? HomeScreen() : LoginScreen();
          },
        ),
      ),
    );
  }
}