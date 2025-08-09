"""Simple LangGraph agent demo."""

from __future__ import annotations

from langchain_core.messages import HumanMessage
from langchain_openai import ChatOpenAI
from langgraph.prebuilt import create_react_agent


def run_agent(prompt: str) -> str:
    """Run the LangGraph agent with a single prompt.

    Parameters
    ----------
    prompt: str
        The user input to send to the agent.

    Returns
    -------
    str
        The assistant's final response.
    """
    llm = ChatOpenAI()
    agent = create_react_agent(llm, tools=[])
    result = ""
    for event in agent.stream({"messages": [HumanMessage(content=prompt)]}, stream_mode="values"):
        result = event["messages"][-1].content
    return result


if __name__ == "__main__":
    print(run_agent("Hello, LangGraph!"))
