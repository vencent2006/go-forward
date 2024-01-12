import Mock from "mockjs"
// 内存模拟数据
const arr = []
for (let i = 0; i < 10; i++) {
  arr.push({
    id: Mock.mock("@id"),
    name: Mock.mock("@cname"),
    place: Mock.mock("@county(true)"),
  })
}
export default [
  {
    url: "/list",
    method: "get",
    response: () => {
      return arr
    },
  },
  {
    url: "/del/:id",
    method: "delete",
    response: (req) => {
      const index = arr.findIndex((item) => item.id === req.query.id)
      if (index > -1) {
        arr.splice(index, 1)
        return { success: true }
      } else {
        return { success: false }
      }
    },
  },
  {
    url: "/edit/:id",
    method: "patch",
    response: ({ query, body }) => {
      const item = arr.find((item) => item.id === query.id)
      if (item) {
        item.name = body.name
        item.place = body.place
        return { success: true }
      } else {
        return { success: false }
      }
    },
  },
]
