package io.dongtai.range.grpc.controller;

import io.dongtai.range.grpc.gencode.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping(value = "/grpc")
public class SendController {
    private final GRPCServiceGrpc.GRPCServiceBlockingStub grpcServiceBlockingStub;

    @Autowired
    public SendController(GRPCServiceGrpc.GRPCServiceBlockingStub grpcServiceBlockingStub) {
        this.grpcServiceBlockingStub = grpcServiceBlockingStub;
    }

    @GetMapping("/send")
    public String send(@RequestParam("text") String text) {
        if (text.isEmpty()){
            return "text can not be empty";
        }
        Response response = this.grpcServiceBlockingStub.send(Request.newBuilder()
                .setText(text)
                .build());
        return response.getMessage();
    }

    private static class SendReq {
        private String text;

        public String getText() {
            return text;
        }
    }

    @PostMapping(path = "/send", consumes = "application/json")
    public String send(@RequestBody SendReq req) {
        if (req.getText().isEmpty()){
            return "text can not be empty";
        }
        Response response = this.grpcServiceBlockingStub.send(Request.newBuilder()
                .setText(req.getText())
                .build());
        return response.getMessage();
    }
}
