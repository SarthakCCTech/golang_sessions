# Goals Todo API - User Stories

## US-01: Create Goal
As a user, I want to create a goal so that I can record what I want to achieve.

Acceptance criteria:
- Accept title, description, status, and target date.
- Title is required.
- Default status is `not_started`.
- Return 201 when created.
- Automatically generate a unique ID for each goal.
- Set the creation timestamp automatically.

## US-02: Get All Goals
As a user, I want to view all goals so that I can track my progress.

Acceptance criteria:
- Return all goals.
- Return an empty list if no goals exist.
- Return 200 on success.
- Return goals in a consistent order, such as newest first.

## US-03: Get Goal by ID
As a user, I want to view a goal by ID so that I can see its details.

Acceptance criteria:
- Return the goal if it exists.
- Return 404 if it does not exist.
- Return 400 if the ID is invalid.
- Include all goal fields in the response.
- Return the correct goal without exposing other goals' data.

## US-04: Update Goal
As a user, I want to update a goal so that I can change its details or status.

Acceptance criteria:
- Update title, description, status, and target date.
- Return 200 on success.
- Return 404 if the goal does not exist.
- Reject an empty or whitespace-only title.
- Accept only `not_started`, `in_progress`, or `completed` as valid statuses.
- Preserve the existing goal if the update request contains invalid data.

## US-05: Delete Goal
As a user, I want to delete a goal so that I can remove goals I no longer need.

Acceptance criteria:
- Delete the goal using its ID.
- Return 204 on success.
- Return 404 if the goal does not exist.
- Return 400 if the ID is invalid.
- Ensure the deleted goal no longer appears in the goal list.


