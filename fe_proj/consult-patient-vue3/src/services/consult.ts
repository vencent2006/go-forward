import type {
  ConsultOrderPreData,
  ConsultOrderPreParams,
  DoctorPage,
  FollowType,
  Image,
  KnowledgePage,
  KnowledgeParams,
  PageParams,
  TopDep,
} from '@/types/consult'
import { request } from '@/utils/request'

// 文章列表
export const getKnowledgePage = (params: KnowledgeParams) =>
  request<KnowledgePage>('patient/home/knowledge', 'GET', params)

// 医生列表
export const getDoctorPage = (params: PageParams) =>
  request<DoctorPage>('home/page/doc', 'GET', params)

// 关注/取消关注
export const followOrUnfollow = (id: string, type: FollowType = 'doc') =>
  request('like', 'POST', { id, type })

// 全部科室
export const getAllDep = () => request<TopDep[]>('dep/all')

// 上传图片
export const uploadImage = (file: File) => {
  const fd = new FormData()
  fd.append('file', file)
  return request<Image>('upload', 'POST', fd)
}

// 获取预支付信息
export const getConsultOrderPre = (params: ConsultOrderPreParams) =>
  request<ConsultOrderPreData>('patient/consult/order/pre', 'GET', params)
