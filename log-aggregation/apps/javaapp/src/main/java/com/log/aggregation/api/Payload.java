package com.log.aggregation.api;

import com.fasterxml.jackson.annotation.JsonProperty;
import javax.ws.rs.core.Response.Status;

public class Payload {
    private int code;

    private String message;

    public Payload() {
    }

    public Payload(int code, String message) {
        this.code = code;
        this.message = message;
    }

    public Payload(Status status) {
        this.code = status.getStatusCode();
        this.message = status.getReasonPhrase();
    }

    @JsonProperty
    public long getCode() {
        return code;
    }

    @JsonProperty
    public String getMessage() {
        return message;
    }
}