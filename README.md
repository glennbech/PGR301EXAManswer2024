In case it is needed here:
Kandidatnr: 2

# Blockchain Busters- Kontinuasjonsoppgave PGR301 2023/2024

 The discussion task source is based on the course material specifically the pptx/pdf slides and additional links will be provided under as "ref: ".
## Task 1A: 
Check github workflow + `Build.yml`

## Task 1B:
To ensure that pull requests to protected branches like main meet certain criteria before they can be merged (such as the code compiling successfully and all tests passing), you can use GitHub's branch protection rules in combination with Continuous Integration (CI) workflows:

1. Go to your GitHub repository and click on "Settings".

2. In the left sidebar, click on "Branches" under the "Code and automation" section.

3. Click on "Add rule" to create a new branch protection rule.

4. In the "Branch name pattern" field, enter main (or your target branch name).

5. Under "Protect matching branches", you can configure various settings. To require the CI checks to pass before merging, check the following options:

* Require status checks to pass before merging: This option enables required status checks.
* Require branches to be up-to-date before merging: This ensures the branch is not out-of-date with the base branch.
* In the "Status checks found in the last week for this repository" section, select the CI workflows you want to require.
* If your workflow is named "Continuous Integration", "Build" or "Build-and-test" you should see it listed there. Running them at least once will be needed for them to show up
    
6. Optionally, you can enable other protections like "Require linear history", "Include administrators", "Require pull request reviews before merging", etc.

7. Click "Create" or "Save changes" to apply the branch protection rule.
## Task 1C:

### Advantages

#### 1. Improved Code Quality: 
With multiple reviewers, the likelihood of catching bugs, logical errors, or issues with code standards increases. Reviewers can offer different perspectives and insights, leading to more robust and maintainable code.

#### 2. Knowledge Sharing: 
The review process facilitates knowledge sharing among team members. Reviewers become familiar with parts of the codebase they might not directly work on, leading to a more cohesive and informed team.

#### 3. Consistency: 
Having multiple approvals helps maintain consistency in code style, architecture, and development practices across the project, as more than one person is ensuring that contributions adhere to the project's standards.

#### 4. Security: 
Multiple reviews can enhance security by ensuring that changes are scrutinized for potential security flaws. This is particularly important in projects with significant security implications.

#### 5. Reduction of "Bus Factor":
The "bus factor" is a measure of the risk associated with the information and capabilities concentrated in a few individuals. Requiring multiple approvers helps mitigate this risk by ensuring that more team members are involved in and understand the codebase.
ref: https://deviq.com/terms/bus-factor
### Disadvantages

#### 1. Slower Development Process: 
Waiting for multiple approvals can slow down the development process, particularly in fast-paced environments or when reviewers are not promptly available.

#### 2. Potential for Bottlenecks:
If the pool of available reviewers is small or if certain individuals are frequently required for reviews due to their expertise, bottlenecks can occur, delaying important changes.

#### 3. Increased Overhead for Small Changes: 
For trivial or non-impactful changes, requiring multiple approvals might be overkill, adding unnecessary overhead and frustration.

#### 4. Merge Conflicts: 
In a fast-moving codebase, waiting for multiple approvals can increase the likelihood of merge conflicts, as the branch being reviewed may fall behind the main branch.

### Balancing the Approach

#### 1. Differentiating Change Types: 
Applying different rules for different types of changes (e.g., critical vs. trivial changes) can help balance speed and thoroughness.

#### 2. Automation: 
Using automated checks for code style, testing, and security can reduce the manual review burden and speed up the process.

#### 3. Clear Guidelines: 
Providing clear, concise guidelines for both contributors and reviewers can streamline the review process and reduce bottlenecks.

#### 4. Asynchronous Reviews: 
Encouraging asynchronous reviews can help minimize delays, allowing reviewers to provide feedback at their convenience.

#### 5. Reviewer Rotation: 
Rotating reviewers can help prevent bottlenecks, distribute workload evenly, and foster broader team knowledge.


## Task 2A: 
Check `infra/*`

## Task 2B:
1. Write in cmd path **/infra:
```shell
terraform init
```
Before you do the step below please check "option 1:" and "option 2:"
2.
```shell
terraform plan
```
3. 
```shell
terraform apply
```
##### Option 1:
Use the -var option to set variables directly on the command line:
1. Replace yourUsername + yourPassword with your Dockerhub username and password.
2. Optional: also add -var="repository_name=yourRepositoryName" or else it will default to "nbx"
```shell
terraform plan -var="dockerhub_username=yourUsername" -var="dockerhub_password=yourPassword"
```
##### Option 2:
Terraform will automatically use environment variables prefixed with TF_VAR_ to populate corresponding Terraform variables.
		
Go to environment variables and set two/three variables:

1.
        Variable = TF_VAR_dockerhub_username
        Value    = your dockerhub username
2.
        Variable = TF_VAR_dockerhub_password
        Value    = your dockerhub password
3.
        Variable = TF_VAR_repository_name
        Value    = default will be "nbx" if not changed
Go to step 3.

## Task 2C:
### Why you get an error:
If you delete the `terraform.tfstate` file and then run again,
```shell 
terraform apply
``` 
Terraform loses its knowledge about what resources it's managing and their current state. So, when Terraform tries to create the Docker Hub repository "nbx" again, it finds that the repository already exists on Docker Hub, leading to an error because Docker Hub repository names must be unique.

The error message you're likely to see will indicate that the repository already exists, and Terraform is attempting to create a resource that conflicts with an existing one.

### How to resolve this issue:
#### 1. Import the existing resource:
You can use the terraform import command to import the existing Docker Hub repository into your Terraform state. This would look something like:
```shell 
terraform import dockerhub_repository.project user/nbx
``` 
"user" should be replaced with the actual docker hub username that is being used.  This command tells Terraform that the existing Docker Hub repository "nbx" should be managed as the resource named "project" in your Terraform configuration.
#### 2. Delete the existing resource:
If it's acceptable to delete the existing repository, you can manually remove it from Docker Hub and then run terraform apply again. This approach is straightforward but might not be suitable if the repository contains important images or tags.
#### 3. Change the repository name:
If neither of the above solutions is suitable, you can change the repository name in your Terraform configuration to something that doesn't already exist on Docker Hub. Update the default value of the repository_name variable in your variables.tf file and run terraform apply again.

## Task 3:
When the sensor has to use the GitHub Actions workflow with their Docker Hub account on their fork of the repository, they will need to make some adjustments and ensure that certain prerequisites are met:
#### 1. Fork the Repository

The sensor should start by forking the repository to their GitHub account. This creates a personal copy of the repository where they can make changes.
#### 2. Configure Docker Hub Credentials as Secrets

To push images to their Docker Hub account, the sensor needs to store their Docker Hub username and password/token as secrets in their forked GitHub repository:

* Go to the forked repository on GitHub.
* Navigate to "Settings" > "Secrets" > "Actions".
* Click on "New repository secret".
* Add a secret named DOCKER_USERNAME and set its value to their Docker Hub username.
* Add a secret named DOCKER_TOKEN (or DOCKER_PASSWORD, depending on what has been used in the workflow file) and set its value to their Docker Hub access token or password.
* Add a secret named REPOSITORY_NAME and set its value to (e.g., "nbx" or some other preferred name)

Docker Hub access tokens can be generated from the Docker Hub settings under "Security". It's recommended to use a token instead of a password for security reasons.
ref: https://docs.docker.com/security/for-developers/access-tokens/

#### 3. Test the Workflow

To test if the changes work as expected, the sensor can push a commit to a branch that triggers the workflow (e.g., main). They should then check the "Actions" tab in their GitHub repository to see if the workflow runs successfully and pushes the Docker image to their Docker Hub account.

### "Hvilke docker kommando kan sensor bruke for å kjøre ditt container image på sin maskin?""
### To run my Image
```sh
docker run -d --name students_container nbx:latest
```
