#!/usr/bin/env python3


import os
import sys
import openai

openai.api_key = os.getenv("OPENAI_API_KEY")

prompt = sys.argv[1]

response = openai.Completion.create(
  model="text-davinci-003",
  prompt=prompt,
  temperature=0,
  max_tokens=4000
)

print(response['choices'][0]['text'])

