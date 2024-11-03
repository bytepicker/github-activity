package main

import "time"

// BasicEventProperties represents the common fields across different event types.
type BasicEventProperties struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		ID           int    `json:"id"`
		Login        string `json:"login"`
		DisplayLogin string `json:"display_login"`
		GravatarID   string `json:"gravatar_id"`
		URL          string `json:"url"`
		AvatarURL    string `json:"avatar_url"`
	} `json:"actor"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
	Org       struct {
		ID         int    `json:"id"`
		Login      string `json:"login"`
		GravatarId string `json:"gravatar_id"`
		URL        string `json:"url"`
		AvatarURL  string `json:"avatar_url"`
	} `json:"org"`
}

// CommitCommentEvent represents a commit comment event.
type CommitCommentEvent struct {
	BasicEventProperties
	Payload struct {
		Action  string
		Comment struct {
			ID   int    `json:"id"`
			Body string `json:"body"`
			URL  string `json:"url"`
		} `json:"comment"`
	} `json:"payload"`
}

// CreateEvent represents a repository, branch, or tag creation event.
type CreateEvent struct {
	BasicEventProperties
	Payload struct {
		Ref          string `json:"ref"`
		RefType      string `json:"ref_type"`
		MasterBranch string `json:"master_branch"`
		Description  string `json:"description"`
		PusherType   string `json:"pusher_type"`
	} `json:"payload"`
}

// DeleteEvent represents a repository, branch, or tag deletion event.
type DeleteEvent struct {
	BasicEventProperties
	Payload struct {
		Ref     string `json:"ref"`
		RefType string `json:"ref_type"`
	} `json:"payload"`
}

// ForkEvent represents a repository fork event.
type ForkEvent struct {
	BasicEventProperties
	Payload struct {
		Forkee struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			Private  bool   `json:"private"`
			Owner    struct {
				Login     string `json:"login"`
				ID        int    `json:"id"`
				AvatarURL string `json:"avatar_url"`
				URL       string `json:"url"`
			} `json:"owner"`
			HTMLURL string `json:"html_url"`
			Fork    bool   `json:"fork"`
			URL     string `json:"url"`
		} `json:"forkee"`
	} `json:"payload"`
}

// GollumEvent represents a wiki page creation or update event.
type GollumEvent struct {
	BasicEventProperties
	Payload struct {
		Pages []struct {
			PageName string `json:"page_name"`
			Title    string `json:"title"`
			Action   string `json:"action"`
			SHA      string `json:"sha"`
			HTMLURL  string `json:"html_url"`
		} `json:"pages"`
	} `json:"payload"`
}

// IssueCommentEvent represents a comment on an issue or pull request event.
type IssueCommentEvent struct {
	BasicEventProperties
	Payload struct {
		Action string `json:"action"`
		Issue  struct {
			ID     int    `json:"id"`
			Number int    `json:"number"`
			Title  string `json:"title"`
			URL    string `json:"url"`
		} `json:"issue"`
		Comment struct {
			ID   int    `json:"id"`
			Body string `json:"body"`
			URL  string `json:"url"`
		} `json:"comment"`
	} `json:"payload"`
}

// IssuesEvent represents an issue-related event.
type IssuesEvent struct {
	BasicEventProperties
	Payload struct {
		Action string `json:"action"`
		Issue  struct {
			ID     int    `json:"id"`
			Number int    `json:"number"`
			Title  string `json:"title"`
			URL    string `json:"url"`
		} `json:"issue"`
	} `json:"payload"`
}

// MemberEvent represents an event where a user is added to a repository.
type MemberEvent struct {
	BasicEventProperties
	Payload struct {
		Action string `json:"action"`
		Member struct {
			Login string `json:"login"`
			ID    int    `json:"id"`
			URL   string `json:"url"`
		} `json:"member"`
	} `json:"payload"`
}

// PublicEvent represents an event where a repository is made public.
type PublicEvent struct {
	BasicEventProperties
	Payload struct{} `json:"payload"`
}

// PullRequestEvent represents an event related to a pull request.
type PullRequestEvent struct {
	BasicEventProperties
	Payload struct {
		Action      string `json:"action"`
		Number      int    `json:"number"`
		PullRequest struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			URL   string `json:"url"`
			Head  struct {
				Ref string `json:"ref"`
				SHA string `json:"sha"`
			} `json:"head"`
			Base struct {
				Ref string `json:"ref"`
				SHA string `json:"sha"`
			} `json:"base"`
		} `json:"pull_request"`
	} `json:"payload"`
}

// PullRequestReviewEvent represents an event where a pull request is reviewed.
type PullRequestReviewEvent struct {
	BasicEventProperties
	Payload struct {
		Action string `json:"action"`
		Review struct {
			ID   int    `json:"id"`
			Body string `json:"body"`
			URL  string `json:"url"`
		} `json:"review"`
		PullRequest struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			URL   string `json:"url"`
		} `json:"pull_request"`
	} `json:"payload"`
}

// PullRequestReviewCommentEvent represents an event for a comment on a pull request review.
type PullRequestReviewCommentEvent struct {
	BasicEventProperties
	Payload struct {
		Action  string `json:"action"`
		Comment struct {
			ID   int    `json:"id"`
			Body string `json:"body"`
			URL  string `json:"url"`
		} `json:"comment"`
		PullRequest struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			URL   string `json:"url"`
		} `json:"pull_request"`
	} `json:"payload"`
}

// PullRequestReviewThreadEvent represents an event where a review thread on a pull request is created.
type PullRequestReviewThreadEvent struct {
	BasicEventProperties
	Payload struct {
		Action string `json:"action"`
		Thread struct {
			ID   int    `json:"id"`
			Body string `json:"body"`
			URL  string `json:"url"`
		} `json:"thread"`
	} `json:"payload"`
}

// PushEvent represents a push to a repository.
type PushEvent struct {
	BasicEventProperties
	Payload struct {
		PushID       int64  `json:"push_id"`
		Size         int    `json:"size"`
		RepositoryID int    `json:"repository_id"`
		DistinctSize int    `json:"distinct_size"`
		Ref          string `json:"ref"`
		Head         string `json:"head"`
		Before       string `json:"before"`
		Commits      []struct {
			Sha    string `json:"sha"`
			Author struct {
				Email string `json:"email"`
				Name  string `json:"name"`
			} `json:"author"`
			Message  string `json:"message"`
			Distinct bool   `json:"distinct"`
			URL      string `json:"url"`
		} `json:"commits"`
	} `json:"payload"`
}

// ReleaseEvent represents a release creation event.
type ReleaseEvent struct {
	BasicEventProperties
	Payload struct {
		Action  string `json:"action"`
		Release struct {
			ID      int    `json:"id"`
			TagName string `json:"tag_name"`
			URL     string `json:"url"`
		} `json:"release"`
	} `json:"payload"`
}

// SponsorshipEvent represents a sponsorship-related event.
type SponsorshipEvent struct {
	BasicEventProperties
	Payload struct {
		Action      string `json:"action"`
		Sponsorship struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"sponsorship"`
	} `json:"payload"`
}

// WatchEvent represents a user watching a repository.
type WatchEvent struct {
	BasicEventProperties
	Payload struct {
		Action string `json:"action"`
	} `json:"payload"`
}
