# Contributing to prOvalada 🏉

Thank you for your interest in contributing to prOvalada! This document provides guidelines and instructions for contributing to the project.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Testing](#testing)
- [Code Style](#code-style)
- [Submitting Changes](#submitting-changes)
- [Reporting Issues](#reporting-issues)

## 📜 Code of Conduct

This project adheres to a code of conduct based on respect, inclusivity, and collaboration. Please be kind and respectful to all community members.

## 🚀 Getting Started

### Prerequisites

- Python 3.8 or higher
- Git
- Docker (optional, but recommended)

### Fork and Clone

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/your-username/prOvalada.git
   cd prOvalada
   ```

## 🛠️ Development Setup

### Option 1: Local Development

```bash
# Create virtual environment
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate

# Install dependencies
pip install -r requirements.txt
pip install -r requirements-dev.txt

# Install pre-commit hooks
pre-commit install
```

### Option 2: Docker Development

```bash
# Start development environment
docker-compose up -d

# Access the container
docker-compose exec api bash
```

## 🔧 Making Changes

### Branch Naming Convention

Use descriptive branch names with prefixes:

- `feature/` - New features
- `fix/` - Bug fixes
- `docs/` - Documentation changes
- `refactor/` - Code refactoring
- `test/` - Test improvements
- `chore/` - Maintenance tasks

Examples:
- `feature/user-authentication`
- `fix/scraper-timeout-issue`
- `docs/api-documentation-update`

### Commit Messages

Follow the conventional commit format:

```
type(scope): description

[optional body]

[optional footer]
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `style`: Code style changes
- `refactor`: Code refactoring
- `test`: Test changes
- `chore`: Maintenance

Examples:
```
feat(api): add rugby match endpoint
fix(scraper): handle timeout errors gracefully
docs(readme): update installation instructions
```

## 🧪 Testing

### Running Tests

```bash
# Run all tests
pytest

# Run with coverage
pytest --cov=app --cov-report=html

# Run specific test categories
pytest -m unit        # Unit tests only
pytest -m integration # Integration tests only
```

### Writing Tests

- Write tests for all new features and bug fixes
- Maintain or improve code coverage
- Use descriptive test names
- Follow the AAA pattern (Arrange, Act, Assert)

Example test structure:
```python
def test_rugby_match_endpoint_returns_valid_data():
    # Arrange
    client = TestClient(app)
    match_id = "12345"
    
    # Act
    response = client.get(f"/matches/{match_id}")
    
    # Assert
    assert response.status_code == 200
    assert "match_date" in response.json()
```

## 🎨 Code Style

### Automated Formatting

This project uses automated code formatting tools:

```bash
# Format code
black .
isort .

# Check formatting
black --check .
isort --check-only .
```

### Linting

```bash
# Run linting
flake8 .
mypy .

# Security checks
bandit -r app/
```

### Style Guidelines

- Follow PEP 8 Python style guide
- Use type hints for function parameters and return values
- Write docstrings for modules, classes, and functions
- Keep line length to 88 characters (Black default)
- Use meaningful variable and function names

## 📤 Submitting Changes

### Pull Request Process

1. **Create a branch** from `main` for your changes
2. **Make your changes** following the guidelines above
3. **Write or update tests** as needed
4. **Run the full test suite** and ensure all tests pass
5. **Run code quality checks** and fix any issues
6. **Update documentation** if needed
7. **Push your branch** to your fork
8. **Create a Pull Request** to the main repository

### Pull Request Template

When creating a PR, please fill out the provided template with:

- Description of changes
- Type of change (bug fix, feature, etc.)
- Related issues
- Testing performed
- Screenshots (if applicable)
- Checklist confirmation

### Review Process

- All PRs require at least one review
- CI pipeline must pass
- Code coverage should not decrease
- Address review feedback promptly
- Keep PRs focused and reasonably sized

## 🐛 Reporting Issues

### Bug Reports

Use the bug report template and include:

- Clear description of the issue
- Steps to reproduce
- Expected vs actual behavior
- Environment details
- Error messages/logs

### Feature Requests

Use the feature request template and include:

- Problem description
- Proposed solution
- Use cases
- Implementation considerations

### Security Issues

For security vulnerabilities, please use GitHub's security advisory feature instead of public issues.

## 🎯 Development Guidelines

### API Development

- Follow RESTful conventions
- Use appropriate HTTP status codes
- Include comprehensive error handling
- Document endpoints with OpenAPI/Swagger
- Implement proper validation with Pydantic

### Database Changes

- Use Alembic migrations for schema changes
- Write both upgrade and downgrade migrations
- Test migrations on sample data
- Include migration in your PR description

### Performance Considerations

- Profile performance-critical code
- Use async/await for I/O operations
- Implement caching where appropriate
- Monitor database query performance

## 📚 Resources

- [FastAPI Documentation](https://fastapi.tiangolo.com/)
- [pytest Documentation](https://docs.pytest.org/)
- [GitHub Flow](https://guides.github.com/introduction/flow/)
- [Semantic Versioning](https://semver.org/)

## ❓ Questions?

- Check existing [GitHub Discussions](https://github.com/facusaldaa/prOvalada/discussions)
- Create a new discussion for questions
- Join our development discussions

Thank you for contributing to prOvalada! 🙏