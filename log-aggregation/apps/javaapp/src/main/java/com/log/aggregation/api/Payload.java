package com.log.aggregation.api;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Payload {
    private int code;

    private String description;

    public Payload() {
    }

    public Payload(int code, String description) {
        this.code = code;
        this.description = description;
    }

    @JsonProperty
    public long getCode() {
        return code;
    }

    @JsonProperty
    public String getDescription() {
        return description;
    }
}