# GitHub Activity Viewer

Implementing [roadmap exapmle project](https://roadmap.sh/projects/github-user-activity)

This is a CLI tool that fetches and displays recent public events for a GitHub user. It uses the GitHub API to retrieve events and organizes them by event type (e.g., push events, issue comments, repository watches). The tool provides a simple overview of activity with details specific to each event type.

## Features

- Retrieves and displays recent events for a specified GitHub username.
- Supports detailed output for various event types, such as:
  - **PushEvent**: Outputs commit messages and authors.
  - **WatchEvent**: Indicates the action when a user stars a repository.
  - **ForkEvent**: Shows the original and forked repository names.
  - **IssueCommentEvent**: Displays the issue number and comment content.
  - **CreateEvent**: Shows the creation of branches, repositories, or tags.
  - **PullRequestEvent**: Shows actions like opening, closing, or merging pull requests.
  - **DeleteEvent**: Indicates deletions of branches or tags.
  - Additional events can be easily added.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or newer)
- A working internet connection (to fetch events from the GitHub API)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/github-activity-viewer.git
   cd github-activity-viewer
   ```

2. Build the project:

   ```bash
   go build -o github-activity
   ```

## Usage

1. Run the binary with a GitHub username as the argument:

   ```bash
   ./github-activity <username>
   ```

   For example:

   ```bash
   ./github-activity octocat
   ```

2. The output will show a list of recent events for the specified user with details based on event type.

## Example Output

```text
- PushEvent in repo octocat/Hello-World by octocat
  Commit: Initial commit by octocat
- WatchEvent: starred on repo octocat/Hello-World
- Forked repo octocat/Hello-World to octocat/Hello-World-fork
- IssueCommentEvent on issue #1: Great work on this issue!
```

## Code Structure

The main components of this project are as follows:

- **Struct Definitions**: Defines structs for each GitHub event type (e.g., `PushEvent`, `WatchEvent`, `ForkEvent`, etc.) with corresponding JSON mappings.
- **Dynamic Unmarshaling**: Parses JSON responses based on event `Type` and unmarshals them into the corresponding structs.
- **Event Handlers**: Switch cases for each event type in `handleEvent` provide specific output logic for each event.

## Adding New Event Types

To add a new event type:
1. Define a struct for the event type (e.g., `CreateEvent`).
2. Add a new case in the `handleEvent` function to handle it specifically.
3. Unmarshal the JSON and display relevant details.

## Notes

- **Rate Limiting**: GitHub API has rate limits for unauthenticated requests. If you hit the limit, try again after an hour or [authenticate requests](https://docs.github.com/en/rest/overview/resources-in-the-rest-api#rate-limiting).
- **Error Handling**: Errors during JSON unmarshaling are logged but won't stop the program from processing other events.

## Contributing

Feel free to submit issues, fork the repository, and send pull requests. Contributions are welcome!

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for more information.