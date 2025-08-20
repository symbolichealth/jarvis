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

### Chat Interface (CLI)
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

### Web Interface
To start the web server:
```bash
$ go run ./cmd/server
```

To start the client:
```bash
$ cd apps/jarvis-webapp && yarn dev
```
Then open your browser to the React webapp at `http://localhost:5173` (see `apps/jarvis-webapp/` for details).

The web server runs on port 7070 and provides:
- `POST /chat` - Chat endpoint for the webapp
- `GET /health` - Health check endpoint

## Docker Setup

To run with PostgreSQL and Temporal using Docker:

1. **Create environment file:**
   ```bash
   cp .env.example .env  # You'll need to create this file manually
   ```

2. **Start services:**
   ```bash
   docker-compose up -d
   ```

3. **Access services:**
   - PostgreSQL: `localhost:4432` (jarvis/jarvis_password) - hosts both `jarvis` and `temporal/temporal_visibility` databases
   - Temporal UI: http://localhost:7080
   - Temporal gRPC: `localhost:6233`

4. **Stop services:**
   ```bash
   docker-compose down
   ```

For detailed Docker setup instructions, see `README-Docker.md`.
