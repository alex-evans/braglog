# BragLog
Logging tool to capture engineering daily / weekly / monthly / yearly notes

## Idea Commands
`brag`
-> brings up today's note with a fresh template if new for the day

```
Today:
[x] PLFM-98 Finish creating tests and submit PR
[ ] PLFM-120 Explore 

On Going Projects Progress:
- PLFM-98 [[hill-80]] 
- PLFM-129 [[hill-25]]
```
hill tag will create a hill hopper view of the progress of a project or task

`brag view`
-> brings up the week's worth of data in html locally (can scroll to other weeks or specific days)

`brag view {date}`
-> brings up the week's view based on that date

`brag view --project {tag}`
-> brings up the project view (when it was worked on and all the updates)

`brag plan`
-> brings up the week in md and allows you to add tasks for the week. Easy way to plan out the week and if something new comes up you can add it to a future day. This will then on that day prepopulate the brag document for that day.

`brag --mtg 1:1`
-> brings up a week's worth of accomplishments in a way that can be easily forwarded or shared with your boss

`brag --mtg {name}`
-> brings up info on the person in question for meeting notes (good way to indicate if you have a question or comment for someone, these will be pulled from the daily notes based on some sort of formatting)

- this could be expanded for managing capabilities of keeping track of reportees and their goals easily
