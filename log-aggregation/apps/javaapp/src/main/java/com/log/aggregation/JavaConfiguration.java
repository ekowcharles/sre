package com.log.aggregation;

import io.dropwizard.Configuration;
import com.fasterxml.jackson.annotation.JsonProperty;

public class JavaConfiguration extends Configuration {
    private String version = "0.0.0";

    @JsonProperty
    public String getVersion() {
        return this.version;
    }

    @JsonProperty
    public void setVersion(String version) {
        this.version = version;
    }
}
