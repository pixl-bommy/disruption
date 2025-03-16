export interface DailyEventItem {
  id: string
  disruptionId: string
  disruptionName: string
  iconName: string
  color: string
  status?: string
  replacedBy?: string
  replaces?: string
}

export type DailyEventItemList = DailyEventItem[]
