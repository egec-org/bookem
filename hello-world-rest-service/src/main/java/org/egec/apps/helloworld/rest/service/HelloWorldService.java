package org.egec.apps.helloworld.rest.service;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.QueryParam;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;

@Path("hello")
public class HelloWorldService {

	@GET
	@Produces(MediaType.TEXT_PLAIN)
	public Response getHelloString(@QueryParam("name") final String name) {
		return Response.ok(String.format("Hello %s!", name)).build();
	}
}
