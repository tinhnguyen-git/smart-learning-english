import 'package:dio/dio.dart';
import 'package:injectable/injectable.dart';
import 'package:retrofit/retrofit.dart';
import '../models/auth_response_model.dart';

part 'auth_remote_data_source.g.dart';

@RestApi(baseUrl: "http://10.0.2.2:8080") // Android emulator generic local host
abstract class AuthRemoteDataSource {
  factory AuthRemoteDataSource(Dio dio, {String baseUrl}) = _AuthRemoteDataSource;

  @POST("/auth/login")
  Future<AuthResponseModel> login(@Body() Map<String, dynamic> body);

  @POST("/auth/register")
  Future<AuthResponseModel> register(@Body() Map<String, dynamic> body);
}


