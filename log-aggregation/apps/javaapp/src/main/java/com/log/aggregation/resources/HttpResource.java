package com.log.aggregation.resources;

import javax.ws.rs.GET;
import javax.ws.rs.OPTIONS;
import javax.ws.rs.DELETE;
import javax.ws.rs.HEAD;
import javax.ws.rs.POST;
import javax.ws.rs.PATCH;
import javax.ws.rs.PUT;
import javax.ws.rs.Path;
import javax.ws.rs.PathParam;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.Response.Status;

import com.codahale.metrics.annotation.Timed;
import com.codahale.metrics.annotation.ResponseMetered;
import com.codahale.metrics.annotation.ExceptionMetered;
import com.log.aggregation.api.Application;
import com.log.aggregation.api.Payload;

@Path("/")
@Produces(MediaType.APPLICATION_JSON)
public class HttpResource {
    private Application application;

    public HttpResource(String version) {
        this.application = new Application(version);
    }

    @GET
    @Timed
    @ResponseMetered
    public Application index() {
        return application;
    }

    @GET
    @Timed
    @Path("/random")
    @ResponseMetered
    @Produces(MediaType.APPLICATION_JSON)
    public Payload random() {

        
        // return new Payload(200, "description");
    }

    @GET
    @Timed
    @Path("/exception")
    @Produces(MediaType.APPLICATION_JSON)
    @ExceptionMetered()
    public Payload exception() throws Exception {
        throw new Exception("Doing this only for the logs");
    }

    @GET
    @Timed
    @Path("/http/{code}")
    @Produces(MediaType.APPLICATION_JSON)
    public Response http_code(@PathParam("code") int code) {
        Status status = Status.fromStatusCode(code);

        return Response.status(status).type(MediaType.APPLICATION_JSON_TYPE)
                .entity(new Payload(status.getStatusCode(), status.getReasonPhrase())).build();
    }
}