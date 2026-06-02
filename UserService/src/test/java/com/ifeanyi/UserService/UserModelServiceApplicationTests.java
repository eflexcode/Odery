package com.ifeanyi.UserService;

import com.ifeanyi.UserService.entity.User;
import com.ifeanyi.UserService.model.UserModel;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.web.client.RestTemplate;

@SpringBootTest
class UserModelServiceApplicationTests {

	static RestTemplate restTemplate;

	@BeforeAll
    static void loadStuff(){
		restTemplate = new RestTemplate();
	}

	@Test
	void contextLoads() {
	}

	@Test
	void  testCreateUser(){

		UserModel userModel = new UserModel();
		userModel.setName("test");
		userModel.setPassword("123456");
		userModel.setImgUrl("dfdfdfdfdf");
		userModel.setUsername("test");

		HttpEntity<UserModel>  httpEntity = new HttpEntity<>(userModel);

		restTemplate.exchange("jjjj", HttpMethod.POST,httpEntity,User.class);
	}

}
