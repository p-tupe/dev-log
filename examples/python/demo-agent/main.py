import asyncio
# import numpy as np
# import pandas as pd
import os

# from google import genai
from google.genai import types

# from IPython.display import Markdown, HTML, display
from google.adk.sessions import InMemorySessionService
from google.adk.runners import Runner
# from google.genai import types # For creating message Content/Parts

GOOGLE_API_KEY = "YOUR-KEY-HERE"
APP_NAME = "demo_app"
USER_ID = "demo_user"
SESSION_ID = "demo_001"

os.environ["GOOGLE_API_KEY"] = GOOGLE_API_KEY

# Simplified Agent for demostration
# -- 1. Agent Definition --
try:
    from google.adk.agents import Agent
except:
    from google.adk.agents import Agent

demo_agent = Agent(
    name="Demo_Agent",
    description="A simplified demo agent",
    model='gemini-2.5-flash',
    instruction="""
    You are an analyst agent that always answer today's date in London,
    Ontario no matter what question is asked
    """,
)


async def main():
    demo_session_service = InMemorySessionService()

    await demo_session_service.create_session(
        app_name=APP_NAME,
        user_id=USER_ID,
        session_id=SESSION_ID
    )

    print(f"Session created: App='{APP_NAME}', \
            User='{USER_ID}', Session='{SESSION_ID}'")

    demo_runner = Runner(
        agent=demo_agent,
        app_name=APP_NAME,
        session_service=demo_session_service
    )

    print(f"Runner created for agent '{demo_runner.agent.name}'.")

    await call_agent_async("How much is a trip from London to Japan?",
                           runner=demo_runner, user_id=USER_ID,
                           session_id=SESSION_ID)


async def call_agent_async(query: str, runner, user_id, session_id):
    """Sends a query to the agent and prints the final response."""
    print(f"\n>>> User Query: {query}")
    content = types.Content(role='user', parts=[types.Part(text=query)])
    final_response_text = "Agent did not produce a final response."
    async for event in runner.run_async(user_id=user_id,
                                        session_id=session_id,
                                        new_message=content):
        if event.is_final_response():
            if event.content and event.content.parts:
                final_response_text = event.content.parts[0].text
            elif event.actions and event.actions.escalate:
                final_response_text = f"Agent escalated: \
                    {event.error_message or 'No specific message.'}"
        break
    print(f"<<< Agent Response: {final_response_text}")


asyncio.run(main())
