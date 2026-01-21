import 'package:fpdart/fpdart.dart';
import 'package:injectable/injectable.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:dio/dio.dart';

import '../../../../core/error/failures.dart';
import '../../domain/entities/user.dart';
import '../../domain/repositories/auth_repository.dart';
import '../datasources/auth_remote_data_source.dart';

@LazySingleton(as: AuthRepository)
class AuthRepositoryImpl implements AuthRepository {
  final AuthRemoteDataSource remoteDataSource;
  final SharedPreferences sharedPreferences;

  AuthRepositoryImpl(this.remoteDataSource, this.sharedPreferences);

  @override
  Future<Either<Failure, User>> login({required String email, required String password}) async {
    try {
      final response = await remoteDataSource.login({'email': email, 'password': password});
      await sharedPreferences.setString('token', response.token);
      return Right(response.user);
    } on DioException catch (e) {
      // Very basic error handling for now
      return Left(ServerFailure(e.message ?? 'Login failed'));
    } catch (e) {
      return Left(ServerFailure(e.toString()));
    }
  }

  @override
  Future<Either<Failure, User>> register({required String fullName, required String email, required String password}) async {
    try {
      final response = await remoteDataSource.register({
        'full_name': fullName,
        'email': email,
        'password': password,
      });
      await sharedPreferences.setString('token', response.token);
      return Right(response.user);
    } on DioException catch (e) {
      return Left(ServerFailure(e.message ?? 'Registration failed'));
    } catch (e) {
      return Left(ServerFailure(e.toString()));
    }
  }

  @override
  Future<Either<Failure, void>> logout() async {
    try {
       await sharedPreferences.remove('token');
       return const Right(null);
    } catch (e) {
      return Left(CacheFailure(e.toString()));
    }
  }

  @override
  Future<Option<User>> getCurrentUser() async {
     // TODO: Implement actual user fetching from stored token or user profile endpoint
     // For now, if token exists, we might need a stored user or fetch it.
     // Skipping deep implementation for this step to focus on initial login flow.
     final token = sharedPreferences.getString('token');
     if (token != null) {
       // Ideally: Fetch user profile from API using token
     }
     return const None();
  }
}
