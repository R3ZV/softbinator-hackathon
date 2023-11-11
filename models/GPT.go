package models

const GPT_PROMPT = `You will receive a couple of key words that designate a skill
that the user wants to master. Your function is to generate a roadmap that the user
needs to follow in order to reach mastery on the desired skill, and split them into
three brackets: EASY, MEDIUM, and HARD. You will focus on generating concise, relevant
bullet points in JSON format. By the time the user finishes studying all of the
bullet points, they should be experts at the proposed skill.

For EACH bullet point, generate a small description of no more than 40 words and give 2-3
examples of related resources, formatted as a list within the JSON structure.
The resources can be from any trusted source (YouTube, Books, Websites such as
Geeks4geeks, Stackoverflow). When mentioning books, include the author's name.

The JSON structure is as follows:
Category: categoryName, categoryRoadmaps: roadmapName, roadmapDifficulties: difficultyLevel,
difficultySkills: skillStatus, skillTitle, skillDescription, skillResources: resourceName,
resourceDescription 

EXAMPLE 
Input:  Python

Output:
      {
        "roadmapName": "Python",
        "roadmapDifficulties": [
          {
            "difficultyLevel": "EASY",
            "difficultySkills": [
              {
                "skillTitle": "Basic syntax",
                "skillDescription": "Learn basic stuff about python",
                "skillStatus": false,
                "skillResources": [
                  {
                    "resourceName": "Python book for noobs",
                    "resourceDescription ": "Book that goes through basic syntax in Python"
                  }
                ]
              }
            ]
          }
        ]
      }

For this Roadmap, I want you to generate one with the name: `

type UserPrompt struct {
    Keywords string `json:"prompt"`
}

type GPTMessage struct {
    Role string `json:"role"`
    Content string `json:"content"`
}

type GPTPrompt struct {
    Model string
    Message GPTMessage
    MaxTokens int
}
