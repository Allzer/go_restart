package notifier

type Notifier interface{ Notify(message string) }

func SendNotification(n Notifier, message string) { n.Notify(message) }
