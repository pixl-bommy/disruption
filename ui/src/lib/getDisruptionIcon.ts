import { defineAsyncComponent, type Component } from 'vue'
import DefaultIcon from '@/components/icons/disruptions/DefaultDisruptionIcon.vue'

/**
 * Get the disruption icon based on the icon name.
 *
 * Tries to load the icon component from the `@/components/icons/disruptions` directory,
 * where `iconName` is the name of the icon file without the extension `DisruptionIcon.vue`.
 *
 * If the icon component is not found, the default icon component is returned.
 *
 * @param iconName The name of the icon to get.
 * @returns The targeted disruption icon component or the default "DefaultDisruptionIcon"
 *      component.
 */
export function getDisruptionIcon(iconName: string) {
  if (!iconName || typeof iconName !== 'string') return DefaultIcon

  const iconFileName = iconName[0].toUpperCase() + iconName.slice(1)

  return defineAsyncComponent<Component>({
    loader: () => import(`@/components/icons/disruptions/${iconFileName}DisruptionIcon.vue`),
    errorComponent: DefaultIcon,
  })
}
