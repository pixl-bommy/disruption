/**
 * A basic disruption item.
 */
export interface DisruptionItem {
  /** The unique identifier of the disruption item. */
  id: string
  /** The name of the disruption item. */
  name: string
  /** The description of the disruption item. */
  description: string
}

/**
 * A list of disruption items.
 */
export type DisruptionItemList = DisruptionItem[]
