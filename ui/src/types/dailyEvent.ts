/**
 * The daily event type as "optimistic" item.
 */
export type DailyEventItem = DailyEventItemUnsettled | DailyEventItemSettled

/**
 * The daily event is not yet settled.
 */
export interface DailyEventItemUnsettled {
  /** temporary, not final id of the event */
  id: string
  /** the unique id of the related disruption, as defined in the client */
  disruptionId: string
  /** the name of the event, as defined in the client */
  name: string
  /** the icon name of the event, as defined in the client */
  iconName: string
  /** the color of the event, as defined in the client */
  color: string

  /** the status of the event */
  status: DailyEventItemStatus.Failed | DailyEventItemStatus.Queued | DailyEventItemStatus.Sending
}

/**
 * The daily event is settled.
 */
export interface DailyEventItemSettled {
  /** the unique id of the event */
  id: string
  /** the unique id of the related disruption */
  disruptionId: string
  /** the name of the event */
  name: string
  /** the icon name of the event */
  iconName: string
  /** the color of the event */
  color: string

  /** the status of the event */
  status: DailyEventItemStatus.Completed
}

export type DailyEventItemList = DailyEventItem[]

export enum DailyEventItemStatus {
  /** the event is queued, to be sent to the backend api */
  Queued,
  /** the event is in sending state */
  Sending,
  /** the event is successfully stored at the backend */
  Completed,
  /** the event is failed to store at the backend */
  Failed,
}
