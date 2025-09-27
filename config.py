# Example health check configuration
# This file will be used once the FastAPI application is implemented

HEALTH_CHECK_CONFIG = {
    "endpoint": "/health",
    "checks": {
        "database": True,
        "redis": True,
        "external_apis": True,
    },
    "timeout": 30,  # seconds
    "interval": 30,  # seconds for Docker healthcheck
}

# Prometheus metrics configuration
METRICS_CONFIG = {
    "endpoint": "/metrics",
    "enabled": True,
    "include_default_metrics": True,
}

# Logging configuration
LOGGING_CONFIG = {
    "version": 1,
    "disable_existing_loggers": False,
    "formatters": {
        "default": {
            "format": "[{time:YYYY-MM-DD HH:mm:ss}] {level} | {name}:{function}:{line} - {message}",
            "style": "{",
        },
        "json": {
            "format": '{"timestamp": "{time:YYYY-MM-DD HH:mm:ss}", "level": "{level}", "logger": "{name}", "function": "{function}", "line": {line}, "message": "{message}"}',
            "style": "{",
        },
    },
    "handlers": {
        "default": {
            "formatter": "default",
            "class": "logging.StreamHandler",
            "stream": "ext://sys.stdout",
        },
    },
    "root": {
        "level": "INFO",
        "handlers": ["default"],
    },
}