package com.example.service.account.controllers;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;

@RestController
class PingController {

    @GetMapping("/ping")
    public String ping() {
        return "Hello Docker World - Account Service";
    }
}