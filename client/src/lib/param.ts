import { useRoute } from 'vue-router'
import { computed, type ComputedRef } from 'vue'

const toStringIfArray = (s: string | undefined | readonly string[]) =>
  Array.isArray(s) ? s[0] : s

/**
 * @param paramName - 取得するparam、存在しないものを取得しようとした場合、'undefined'が返るので注意
 */
const useParam = (paramName: string): ComputedRef<string> => {
  const route = useRoute()
  return computed(() => toStringIfArray(route.params[paramName]) ?? 'undefined')
}

export default useParam
