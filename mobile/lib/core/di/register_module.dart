import 'package:dio/dio.dart';
import 'package:injectable/injectable.dart';
import 'package:isar/isar.dart';
import 'package:path_provider/path_provider.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../../features/auth/data/datasources/auth_remote_data_source.dart';

@module
abstract class RegisterModule {
  @lazySingleton
  Dio get dio => Dio();

  @preResolve
  Future<SharedPreferences> get prefs => SharedPreferences.getInstance();

  @lazySingleton
  AuthRemoteDataSource authRemoteDataSource(Dio dio) => AuthRemoteDataSource(dio);

  @preResolve
  Future<Isar> get isar async {
    final dir = await getApplicationDocumentsDirectory();
    return Isar.open(
      [], // Add schemas here later
      directory: dir.path,
    );
  }
}
