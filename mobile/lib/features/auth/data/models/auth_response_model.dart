import 'package:json_annotation/json_annotation.dart';
import '../../domain/entities/user.dart';

part 'auth_response_model.g.dart';

@JsonSerializable()
class AuthResponseModel {
  @JsonKey(name: 'token')
  final String token;
  
  @JsonKey(name: 'user')
  final UserModel user;

  const AuthResponseModel({required this.token, required this.user});

  factory AuthResponseModel.fromJson(Map<String, dynamic> json) => 
      _$AuthResponseModelFromJson(json);
  
  Map<String, dynamic> toJson() => _$AuthResponseModelToJson(this);
}

@JsonSerializable()
class UserModel extends User {
  const UserModel({
    required super.id,
    required super.email,
    @JsonKey(name: 'full_name') required super.fullName,
    @JsonKey(name: 'is_premium') super.isPremium = false,
  });

  factory UserModel.fromJson(Map<String, dynamic> json) => 
      _$UserModelFromJson(json);
  
  Map<String, dynamic> toJson() => _$UserModelToJson(this);
}
