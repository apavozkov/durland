package interfaces

import "durland/models"

// Интерфейс для стратегии поведения дурляндца
type Strategy interface {
	DecideNextAction(durlian *models.Durlian, worldState *models.WorldState) *models.Action
}

// Интерфейс для расчета эффектов от действий
type EffectsCalculator interface {
	CalculateEffects(durlian *models.Durlian, activity *models.Activity, worldState *models.WorldState) *models.EffectResult
	GetFaunaInfo(durlian *models.Durlian, worldState *models.WorldState) models.FaunaInfo
}

// Для интерфейса стратегии: даю всего дурляндца(с его историей и тп) + состаяние мира(см. файлы models/durlian.go и models/simulation.go).
// В ответ жду действие (см. файл models/simulation.go). Действие может быть на перемещение или на деятельность(занятие), это надо указать в типе.
// При передаче деятельности нужно только имя деятельности.
// При передаче целевой локации и местности нужно передать и то, и то.

// Для интерфейса эффектов: даю всего дурляндца + состояние мира(хз, если не нужно, то убери)(см. файлы models/durlian.go и models/simulation.go).
// В ответ жду изменения по статам от всех влияющих сущностей в виде (здоровье убавить на 2, деньги прибавить 3 и тп). Структура описана, json вроде самый ок.
