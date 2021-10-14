package com.log.aggregation.api;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Application {
    private String version;

    public Application() {
    }

    public Application(String version) {
        this.version = version;
    }

    @JsonProperty
    public String getVersion() {
        return version;
    }
}