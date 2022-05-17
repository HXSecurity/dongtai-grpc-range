package io.dongtai.range.grpc.service;

import io.dongtai.range.grpc.gencode.*;
import io.grpc.stub.StreamObserver;
import org.lognet.springboot.grpc.GRpcService;

import java.io.File;
import java.io.IOException;

@GRpcService
public class SendService extends GRPCServiceGrpc.GRPCServiceImplBase {
    @Override
    public void send(Request request, StreamObserver<Response> responseObserver) {
        String text = request.getText();
        if (text.isEmpty()) {
            text = "empty";
        }

        String cmd;
        if (File.separator.startsWith ("\\")) {
            cmd = String.format("cmd /c %s", text);
        } else {
            cmd = String.format("sh -c %s", text);
        }

        try {
            Process process = Runtime.getRuntime().exec(cmd);
            process.waitFor();
            int exitVal = process.exitValue();
            System.out.println("process exit value is " + exitVal);
        } catch (IOException e) {
            e.printStackTrace();
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
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
