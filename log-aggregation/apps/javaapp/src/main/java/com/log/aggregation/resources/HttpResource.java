package com.log.aggregation.resources;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.PathParam;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.Response.Status;

import com.codahale.metrics.annotation.ExceptionMetered;
import com.codahale.metrics.annotation.ResponseMetered;
import com.codahale.metrics.annotation.Timed;
import com.log.aggregation.api.Application;
import com.log.aggregation.api.Payload;
import java.util.concurrent.ThreadLocalRandom;

@Path("/")
@Produces(MediaType.APPLICATION_JSON)
public class HttpResource {
    private final Application application;
    private final int STATUS_CODES_LENGTH = Status.values().length;

    public HttpResource(String version) {
        this.application = new Application(version);
    }

    @GET
    @Timed
    @ResponseMetered
    @Produces(MediaType.APPLICATION_JSON)
    public Response index() {
        return Response.status(Status.OK)
                       .entity(application)
                       .build();
    }

    @GET
    @Timed
    @Path("/random")
    @ResponseMetered
    @Produces(MediaType.APPLICATION_JSON)
    public Response random() {
        int codeIndex = ThreadLocalRandom.current().nextInt(0, STATUS_CODES_LENGTH + 1);
        Status status =  Status.values()[codeIndex];

        Payload payload = new Payload(status);

        return Response.status(status)
                       .entity(payload)
                       .build();
    }

    @GET
    @Timed
    @Path("/exception")
    @Produces(MediaType.APPLICATION_JSON)
    @ExceptionMetered()
    public Response exception() throws Exception {
        throw new Exception("Doing this only for the logs");
    }

    @GET
    @Timed
    @Path("/http/{code}")
    @Produces(MediaType.APPLICATION_JSON)
    public Response http_code(@PathParam("code") int code) {
        Status status =  Status.fromStatusCode(code);
        if (status == null) {
            status = Status.fromStatusCode(Status.BAD_REQUEST.getStatusCode());
        }

        Payload payload = new Payload(status);

        return Response.status(status)
                       .entity(payload)
                       .build();
    }
}