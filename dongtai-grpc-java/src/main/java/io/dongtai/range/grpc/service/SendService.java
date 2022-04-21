package io.dongtai.range.grpc.service;

import io.dongtai.range.grpc.gencode.*;
import io.grpc.stub.StreamObserver;
import org.lognet.springboot.grpc.GRpcService;

@GRpcService
public class SendService extends GRPCServiceGrpc.GRPCServiceImplBase {
    @Override
    public void send(Request request, StreamObserver<Response> responseObserver) {
        String text = request.getText();
        if (text.isEmpty()) {
            text = "empty";
        }

        text = "text: " + text;
        Response response = Response.newBuilder()
                .setMessage(text)
                .setState(200)
                .build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}
