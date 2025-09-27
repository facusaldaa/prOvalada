# prOvalada 🏉

[![CI/CD Pipeline](https://github.com/facusaldaa/prOvalada/actions/workflows/ci-cd.yml/badge.svg)](https://github.com/facusaldaa/prOvalada/actions/workflows/ci-cd.yml)
[![codecov](https://codecov.io/gh/facusaldaa/prOvalada/branch/main/graph/badge.svg)](https://codecov.io/gh/facusaldaa/prOvalada)
[![Python 3.8+](https://img.shields.io/badge/python-3.8+-blue.svg)](https://www.python.org/downloads/)
[![FastAPI](https://img.shields.io/badge/FastAPI-0.104+-00a393.svg)](https://fastapi.tiangolo.com)
[![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)
[![License](https://img.shields.io/github/license/facusaldaa/prOvalada)](https://github.com/facusaldaa/prOvalada/blob/main/LICENSE)

API REST para datos y estadísticas de rugby argentino. Scraping automático de fixtures, resultados y torneos desde ESPN Argentina.

## ✨ Features

- 🏉 **Rugby Data API**: Comprehensive REST API for Argentine rugby statistics
- 🔄 **Automatic Scraping**: Real-time data collection from ESPN Argentina
- ⚡ **Fast & Async**: Built with FastAPI for high performance
- 🐳 **Containerized**: Docker support for easy deployment
- 🔒 **Secure**: Built-in authentication and security features
- 📊 **Monitoring**: Prometheus metrics and health checks
- 🧪 **Well Tested**: Comprehensive test suite with high coverage

## 🚀 Quick Start

### Using Docker (Recommended)

```bash
# Clone the repository
git clone https://github.com/facusaldaa/prOvalada.git
cd prOvalada

# Start with Docker Compose
docker-compose up -d

# The API will be available at http://localhost:8000
```

### Local Development

```bash
# Create virtual environment
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate

# Install dependencies
pip install -r requirements.txt
pip install -r requirements-dev.txt

# Start the development server
uvicorn app.main:app --reload --host 0.0.0.0 --port 8000
```

## 📖 API Documentation

Once the server is running, visit:

- **Interactive API Docs (Swagger UI)**: http://localhost:8000/docs
- **Alternative API Docs (ReDoc)**: http://localhost:8000/redoc
- **OpenAPI JSON**: http://localhost:8000/openapi.json

## 🏗️ Development

### Prerequisites

- Python 3.8+ 
- Docker & Docker Compose (for containerized development)
- Git

### Setup Development Environment

```bash
# Clone and setup
git clone https://github.com/facusaldaa/prOvalada.git
cd prOvalada

# Install development dependencies
pip install -r requirements-dev.txt

# Install pre-commit hooks
pre-commit install

# Run tests
pytest

# Run with coverage
pytest --cov=app --cov-report=html
```

### Code Quality

This project uses several tools to maintain code quality:

- **Black**: Code formatting
- **isort**: Import sorting
- **flake8**: Linting
- **mypy**: Type checking
- **bandit**: Security linting
- **pre-commit**: Git hooks for quality checks

```bash
# Format code
black .
isort .

# Lint code
flake8 .
mypy .

# Security scan
bandit -r .

# Run all pre-commit hooks
pre-commit run --all-files
```

### Testing

```bash
# Run all tests
pytest

# Run with coverage
pytest --cov=app --cov-report=html --cov-report=term

# Run specific test categories
pytest -m unit        # Unit tests only
pytest -m integration # Integration tests only
pytest -m "not slow"  # Exclude slow tests
```

## 🚀 CI/CD Pipeline

This project includes a comprehensive GitHub Actions CI/CD pipeline:

### Pipeline Stages

1. **Code Quality & Security**
   - Code formatting (Black)
   - Import sorting (isort) 
   - Linting (Flake8)
   - Type checking (MyPy)
   - Security scanning (Bandit, Safety)

2. **Testing**
   - Multi-version Python testing (3.8, 3.9, 3.10, 3.11)
   - Coverage reporting
   - Integration with Codecov

3. **Build**
   - Package building
   - Docker image creation
   - Multi-platform support (AMD64, ARM64)

4. **Deployment**
   - Staging deployment (develop branch)
   - Production deployment (main branch)
   - Manual deployment triggers

### Workflow Triggers

- **Pull Requests**: Full CI validation
- **Push to main/develop**: CI + CD pipeline
- **Manual Triggers**: Flexible deployment options
- **Dependabot**: Automated dependency updates

## 🐳 Docker

### Multi-stage Dockerfile

- Optimized for production with minimal image size
- Non-root user for security
- Health checks included
- Multi-platform builds (AMD64/ARM64)

### Docker Compose

```bash
# Development environment
docker-compose up -d

# Production environment
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d

# View logs
docker-compose logs -f api
```

## 📊 Monitoring & Observability

- **Health Checks**: `/health` endpoint
- **Metrics**: Prometheus metrics at `/metrics`
- **Logging**: Structured logging with Loguru
- **Tracing**: Application performance monitoring ready

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

### Development Workflow

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests and quality checks
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

### Code Review Process

- All changes require review via Pull Request
- CI pipeline must pass
- Code coverage should not decrease
- Follow the existing code style and conventions

## 📈 Roadmap

- [ ] **Phase 1**: Basic API structure and scraping
- [ ] **Phase 2**: Authentication and user management
- [ ] **Phase 3**: Real-time data streaming
- [ ] **Phase 4**: Mobile app integration
- [ ] **Phase 5**: Analytics dashboard

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- ESPN Argentina for providing rugby data
- FastAPI community for the excellent framework
- All contributors who help improve this project

## 📞 Support

- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/facusaldaa/prOvalada/issues)
- 💡 **Feature Requests**: [GitHub Issues](https://github.com/facusaldaa/prOvalada/issues)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/facusaldaa/prOvalada/discussions)

---

Made with ❤️ by [Facundo Salda](https://github.com/facusaldaa)
