package models

// Типы эффектов окружения, чтобы не дублировать строковые литералы
const (
	EffectSlesandraDamageChance               = "slesandra_damage_chance"
	EffectSlesandraProductivityBonus          = "slesandra_productivity_bonus"
	EffectSatisfactionCostMultiplier          = "satisfaction_cost_multiplier"
	EffectSisandraFatigueProbability          = "sisandra_fatigue_probability_after_first"
	EffectSisandraProductivityBonusAfterFirst = "sisandra_productivity_bonus_after_first"
	EffectMoneyWipeOnStayProbability          = "money_wipe_on_stay_probability"
	EffectChuchundraProductivityBonus         = "chuchundra_productivity_bonus"
	EffectDrocentsHealthPenaltyPerStep        = "drocents_health_penalty_per_step"
)

// Типы эффектов народов

const (
	EffectMozhoryGulbonitMoneyIncrease             = "mozhory_gulbonit_money_increase"
	EffectMozhoryZumbalitHealthSave                = "mozhory_zumbalit_health_save"
	EffectNishcheboryGulbonitMoneyDecrease         = "nishchebory_gulbonit_money_decrease"
	EffectNishcheboryGulbonitHealthIncrease        = "nishchebory_gulbonit_health_increase"
	EffectSoyevyeZumbalitHealthPenalty             = "soyevye_zumbalit_health_penalty"
	EffectProsvetlennyeShlyamsatHistoryBonus       = "prosvetlennye_shlyamsat_history_bonus"
	EffectDrotsentyGulbonitEfficiencyDecrease      = "drotsenty_gulbonit_efficiency_decrease"
	EffectZheleznoukhiZumbalitSatisfactionImmunity = "zheleznoukhi_zumbalit_satisfaction_immunity"
	EffectZheleznoukhiZumbalitMoneyLossChance      = "zheleznoukhi_zumbalit_money_loss_chance"
)

// Типы эффектов локаций

const (
	EffectBalbesburgSlesandraDamageChance        = "balbesburg_slesandra_damage_chance"
	EffectDolbesburgSlesandraProductivityBonus   = "dolbesburg_slesandra_productivity_bonus"
	EffectDolbesburgSatisfactionCostMultiplier   = "dolbesburg_satisfaction_cost_multiplier"
	EffectKuramaribySisandraFatigueProbability   = "kuramariby_sisandra_fatigue_probability"
	EffectPuntaPelikanaSisandraProductivityBonus = "punta_pelikana_sisandra_productivity_bonus"
	EffectPuntaPelikanaMoneyWipeProbability      = "punta_pelikana_money_wipe_probability"
	EffectShrinavasChuchundraProductivityBonus   = "shrinavas_chuchundra_productivity_bonus"
	EffectHareKirishiDrotsentyHealthPenalty      = "hare_kirishi_drotsenty_health_penalty"
)
