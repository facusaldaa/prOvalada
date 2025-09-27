.PHONY: help install install-dev test lint format clean docker-build docker-run docker-stop

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

install: ## Install production dependencies
	pip install -r requirements.txt

install-dev: ## Install development dependencies
	pip install -r requirements.txt -r requirements-dev.txt
	pre-commit install

test: ## Run tests
	pytest

test-cov: ## Run tests with coverage
	pytest --cov=app --cov-report=html --cov-report=term

lint: ## Run linters
	flake8 .
	mypy .
	bandit -r . -x tests/

format: ## Format code
	black .
	isort .

format-check: ## Check code formatting
	black --check .
	isort --check-only .

security: ## Run security checks
	bandit -r . -x tests/
	safety check

clean: ## Clean up cache and temporary files
	find . -type d -name "__pycache__" -exec rm -rf {} + 2>/dev/null || true
	find . -type f -name "*.pyc" -delete
	find . -type f -name "*.pyo" -delete
	find . -type f -name "*~" -delete
	rm -rf .pytest_cache/
	rm -rf .mypy_cache/
	rm -rf .coverage
	rm -rf htmlcov/
	rm -rf dist/
	rm -rf build/
	rm -rf *.egg-info

docker-build: ## Build Docker image
	docker build -t provalada:latest .

docker-run: ## Run application with Docker Compose
	docker-compose up -d

docker-stop: ## Stop Docker Compose services
	docker-compose down

docker-logs: ## Show Docker Compose logs
	docker-compose logs -f

dev-setup: ## Setup development environment
	python -m venv venv
	. venv/bin/activate && make install-dev
	cp .env.example .env
	@echo "Development environment setup complete!"
	@echo "Activate with: source venv/bin/activate"

requirements-update: ## Update requirements files
	pip-compile --upgrade requirements.in
	pip-compile --upgrade requirements-dev.in