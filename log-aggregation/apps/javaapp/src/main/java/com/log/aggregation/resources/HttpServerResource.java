package com.log.aggregation.resources;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

import com.codahale.metrics.annotation.Timed;
import com.log.aggregation.api.Application;

@Path("/")
@Produces(MediaType.APPLICATION_JSON)
public class HttpServerResource {
    private Application application;

    public HttpServerResource(String version) {
        this.application = new Application(version);
    }

    @GET
    @Timed
    public Application index() {
        return application;
    }
}