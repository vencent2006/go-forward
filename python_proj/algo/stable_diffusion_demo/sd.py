import diffusers
import torch
from diffusers import StableDiffusionPipeline

"""
这个在m1上不能用，主要看sd2.py
"""

print(diffusers.__version__)
PROMPT = "a cute cat"
YOUR_TOKEN = "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"
DEVICE, DTYPE = ("cuda", torch.float16) \
    if torch.cuda.is_available() \
    else ("cpu", torch.bfloat16)

pipe = StableDiffusionPipeline.from_pretrained(
    "CompVis/stable-diffusion-v1-4",
    use_auth_token=YOUR_TOKEN,
    torch_dtype=DTYPE
).to(DEVICE)

with torch.autocast(DEVICE, dtype=DTYPE):
    image = pipe(PROMPT)["sample"][0]

image.save("image.png")
