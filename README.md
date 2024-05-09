# TestGo

TestGo is a Go project showcasing a structured file system following MVC (Model-View-Controller) architecture, utilizing the Fiber web framework, GORM for ORM, and SQLite for database management. The project is dockerized for easy deployment and management.

## Technologies Used

- **Go 1.22**: The backend language used for development.
- **Fiber**: A fast, express-like web framework for Go.
- **GORM**: An ORM library for Go, providing a convenient way to interact with databases.
- **SQLite**: A lightweight, serverless database engine used for data storage.
- **Docker**: Containerization technology for packaging the application and its dependencies.

## Project Structure

The project follows a structured file system adhering to MVC principles, ensuring clean separation of concerns and easy maintainability. Here's a brief overview of the directory structure:

- **/cmd/server**: Contains the main application logic.
- **/pkg**
    - **/controllers**: Controllers responsible for handling HTTP requests and responses.
    - **/models**: Defines database models and interacts with the database using GORM.
    - **/config**: Configuration files for the application.
    - **/database**: Contains database related files.
    - **/docker**: Docker configuration files.
    - **/routes**: Defines application routes.
    - **/middleware**: Middleware functions to be applied to routes.
    - **/utils**: Utility functions used across the application.

## Usage

1. Clone the repository:

    ```bash
    git clone https://github.com/mahnueldev/testgo.git
    ```

2. Navigate to the project directory:

    ```bash
    cd testgo
    ```

3. Build and run the Docker container:

    ```bash
    docker-compose up --build
    ```

4. Access the application at `http://localhost:8080`.

## Contributing

Contributions are welcome! Feel free to open issues or pull requests for any improvements or features you'd like to see added to the project.

## License

This project is licensed under the [MIT License](LICENSE).
