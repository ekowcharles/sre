package com.log.aggregation.health;

import com.codahale.metrics.health.HealthCheck;

public class AppHealthCheck extends HealthCheck {
    public AppHealthCheck() {
    }

    @Override
    protected Result check() throws Exception {
        return Result.healthy();
    }
}