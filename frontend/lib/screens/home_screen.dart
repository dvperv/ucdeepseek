import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../services/auth_service.dart';

class HomeScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final authService = Provider.of<AuthService>(context);

    return Scaffold(
      appBar: AppBar(
        title: Text('Home'),
        actions: [
          IconButton(
            icon: Icon(Icons.logout),
            onPressed: () {
              authService.logout();
            },
          ),
        ],
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('Welcome to UnityComm!'),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: () {
                // TODO: Navigate to chat
              },
              child: Text('Go to Chat'),
            ),
            ElevatedButton(
              onPressed: () {
                // TODO: Navigate to email
              },
              child: Text('Go to Email'),
            ),
          ],
        ),
      ),
    );
  }
}