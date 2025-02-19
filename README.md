# Weather Forecast Application

This is a simple Go application that fetches and displays weather forecast information from the WeatherAPI.

## Prerequisites

- Go 1.23.6 or later
- An API key from [WeatherAPI](https://www.weatherapi.com/)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/fsilvestri/sun.git
    cd sun
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

## Usage

1. Run the application:

    ```sh
    go run main.go [location]
    ```

    Replace `[location]` with the desired location. If no location is provided, it defaults to `florianopolis`.

2. Example:

    ```sh
    go run main.go bombinhas
    ```

## Output

The application will display the current weather and hourly forecast for the specified location. If the chance of rain is 40% or higher, the forecast will be displayed in red.

## Files

- [main.go](http://_vscodecontentref_/0): The main application file.
- [go.mod](http://_vscodecontentref_/1): The Go module file.
- [go.sum](http://_vscodecontentref_/2): The Go dependencies file.
- [response.json](http://_vscodecontentref_/3): A sample response from the WeatherAPI.
