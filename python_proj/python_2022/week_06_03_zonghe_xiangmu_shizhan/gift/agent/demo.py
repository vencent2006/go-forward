from langchain_core.output_parsers import StrOutputParser
from langchain_core.prompts import ChatPromptTemplate
from langchain_openai import ChatOpenAI

format_instructions = """输出应格式化为符合以下JSON结构的JSON实例。
{
    "score": 5,
    "reason": "原因"
}
"""

prompt = ChatPromptTemplate.from_template(
    "{format_instructions}\n这是一条微博评论，帮我判断它的内容是否是正向的，并从1到10进行打分，1代表特别负向，10代表特别正向: {comment}")
model = ChatOpenAI(
    base_url="https://api.deepseek.com/v1",
    api_key=os.getenv("DEEPSEEK_API_KEY"),
    model="deepseek-chat",  # deepseek-V3
    temperature=0.7
)
output_parser = StrOutputParser()

chain = prompt | model | output_parser
res = chain.invoke(
    {
        "format_instructions": format_instructions,
        "comment": "太温暖、感人了，看得我热泪盈眶。今年其实有在这样去做了，30多了第一次解锁游乐场、第一次看演唱会，利用周末时间出去旅游。虽然今年工作更加忙碌，压力也更大，但还是很开心，感觉生活开始有了属于自己的色彩。明年会继续努力的"
    })
print(res)
