package com.log.aggregation;

import com.log.aggregation.resources.HttpResource;
import com.log.aggregation.health.AppHealthCheck;

import io.dropwizard.Application;
import io.dropwizard.setup.Bootstrap;
import io.dropwizard.setup.Environment;
import io.dropwizard.configuration.SubstitutingSourceProvider;
import io.dropwizard.configuration.EnvironmentVariableSubstitutor;

public class JavaApplication extends Application<JavaConfiguration> {

    public static void main(final String[] args) throws Exception {
        new JavaApplication().run(args);
    }

    @Override
    public String getName() {
        return "javaapp";
    }

    @Override
    public void initialize(final Bootstrap<JavaConfiguration> bootstrap) {
        bootstrap.setConfigurationSourceProvider(new SubstitutingSourceProvider(
            bootstrap.getConfigurationSourceProvider(), new EnvironmentVariableSubstitutor(false)));
    }

    @Override
    public void run(final JavaConfiguration configuration, final Environment environment) {
        environment.healthChecks().register("apphealth", new AppHealthCheck());
        environment.jersey().register(new HttpResource(configuration.getVersion()));
    }
}