# Jarvis
Jarvis is an LLM powered personal assistant that can do generally useful things, such as:
* Notion project manager
* Business plan creator
* Deep research that is configurable
* Data miner
* Financial advisor
* Logistics planner
* Diagram creator
* Rubber ducky

The only real requirements are:
1. The language has to be golang (sorry, I've been burned by python)
2. The repo should be plug-n-playable

Anything else is really just a suggestion. Ultimately, users should be able to fork the repo, run `./jarvis ...` and get utility.

## Usage
```
$ go run cmd/jarvis/main.go chat
> What's the population of the United States?

Jarvis: As of **early May 2024**, the estimated population of the United States is around **336 million people**.

The U.S. Census Bureau's "Population Clock" provides a real-time estimate, which is constantly changing due to births, deaths, and net international migration. You can find the most up-to-the-minute figure on their website.
> Show me the countries that have higher population

Jarvis: Based on the current estimated population of the United States (around 336 million), there are currently **two** countries that have a higher population:

1.  **India:** Estimated population over **1.4 billion**
2.  **China:** Estimated population over **1.4 billion**

These two countries have significantly larger populations than the United States. The third most populous country after India and China is the United States itself.
> 
```
