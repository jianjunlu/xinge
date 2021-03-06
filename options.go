package xinge

import (
	"time"
)

//OptionPlatIos 的注释
func OptionPlatIos() PushMsgOption {
	return OptionPlatform(PlatformiOS)
}

//OptionPlatAndroid 的注释
func OptionPlatAndroid() PushMsgOption {
	return OptionPlatform(PlatformAndroid)
}

//OptionPlatform 的注释
func OptionPlatform(p Platform) PushMsgOption {
	return func(r *PushMsg) error {
		r.Platform = p
		if p == PlatformiOS {
			r.Android = nil
			r.IOS = &IOSParams{}
		} else if p == PlatformAndroid {
			r.IOS = nil
			r.Android = &AndroidParams{}
		}
		return nil
	}
}

//OptionEnvProduction 的注释
func OptionEnvProduction() PushMsgOption {
	return func(r *PushMsg) error {
		r.Environment = EnvProd
		return nil
	}
}

//OptionEnvDevelop 的注释
func OptionEnvDevelop() PushMsgOption {
	return func(r *PushMsg) error {
		r.Environment = EnvDev
		return nil
	}
}

//OptionTitle 的注释
func OptionTitle(t string) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Title = t
		if r.Message.IOS != nil {
			if r.Message.IOS.Aps != nil {
				r.Message.IOS.Aps.Alert["title"] = t
			} else {
				r.Message.IOS.Aps = &Aps{
					Alert: map[string]interface{}{"title": t},
				}
			}
		} else {
			r.Message.IOS = &IOSParams{
				Aps: &Aps{
					Alert: map[string]interface{}{"title": t},
				},
			}
		}
		return nil
	}
}

//OptionContent 的注释
func OptionContent(c string) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Content = c
		if r.Message.IOS != nil {
			if r.Message.IOS.Aps != nil {
				r.Message.IOS.Aps.Alert["body"] = c
			} else {
				r.Message.IOS.Aps = &Aps{
					Alert: map[string]interface{}{"body": c},
				}
			}
		} else {
			r.Message.IOS = &IOSParams{
				Aps: &Aps{
					Alert: map[string]interface{}{"body": c},
				},
			}
		}
		return nil
	}
}

//OptionAndroidParams 设置AndroidParams
func OptionAndroidParams(params *AndroidParams) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.Android = params
		return nil
	}
}

//OptionIOSParams 设置IOSParams
func OptionIOSParams(params *IOSParams) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message.IOS = params
		return nil
	}
}

//OptionNID 的注释
func OptionNID(id int) PushMsgOption {
	return func(r *PushMsg) error {
		checkAndroidParams(r)
		r.Message.Android.NID = id
		return nil
	}
}
func checkAndroidParams(r *PushMsg) {
	if r.Message.Android == nil {
		r.Message.Android = DefaultAndroidParams()
	}
}
func checkIOSParams(r *PushMsg) {
	if r.Message.IOS == nil {
		r.Message.IOS = DefaultIOSParams(r.Title, r.Content)
	}
}

//OptionBuilderID 的注释
func OptionBuilderID(id int) PushMsgOption {
	return func(r *PushMsg) error {
		checkAndroidParams(r)
		r.Message.Android.BuilderID = id
		return nil
	}
}

//OptionRing 的注释
func OptionRing(ring int) PushMsgOption {
	return func(r *PushMsg) error {
		checkAndroidParams(r)
		r.Message.Android.Ring = ring
		return nil
	}
}

//OptionRingRaw 的注释
func OptionRingRaw(rr string) PushMsgOption {
	return func(r *PushMsg) error {
		checkAndroidParams(r)
		r.Message.Android.RingRaw = rr
		return nil
	}
}

//OptionVibrate 的注释
func OptionVibrate(v int) PushMsgOption {
	return func(r *PushMsg) error {
		checkAndroidParams(r)
		r.Message.Android.Vibrate = v
		return nil
	}
}

//OptionLights 的注释
func OptionLights(l int) PushMsgOption {
	return func(r *PushMsg) error {
		checkAndroidParams(r)
		r.Message.Android.Lights = l
		return nil
	}
}

//OptionCleanable 的注释
func OptionCleanable(c int) PushMsgOption {
	return func(r *PushMsg) error {
		checkAndroidParams(r)
		r.Message.Android.Cleanable = c
		return nil
	}
}

//OptionIconType 的注释
func OptionIconType(it int) PushMsgOption {
	return func(r *PushMsg) error {
		checkAndroidParams(r)
		r.Message.Android.IconType = it
		return nil
	}
}

//OptionIconRes 的注释
func OptionIconRes(ir string) PushMsgOption {
	return func(r *PushMsg) error {
		checkAndroidParams(r)
		r.Message.Android.IconRes = ir
		return nil
	}
}

//OptionStyleID 的注释
func OptionStyleID(s int) PushMsgOption {
	return func(r *PushMsg) error {
		checkAndroidParams(r)
		r.Message.Android.StyleID = s
		return nil
	}
}

//OptionSmallIcon 的注释
func OptionSmallIcon(si int) PushMsgOption {
	return func(r *PushMsg) error {
		checkAndroidParams(r)
		r.Message.Android.SmallIcon = si
		return nil
	}
}

//OptionAddAction 的注释
func OptionAddAction(k string, v interface{}) PushMsgOption {
	return func(r *PushMsg) error {
		checkAndroidParams(r)
		if r.Message.Android.Action == nil {
			r.Message.Android.Action = map[string]interface{}{k: v}
		} else {
			r.Message.Android.Action[k] = v
		}
		return nil
	}
}

//OptionCustomContent 的注释
func OptionCustomContent(ct map[string]string) PushMsgOption {
	return func(r *PushMsg) error {
		if r.Platform == PlatformAndroid {
			checkAndroidParams(r)
			r.Message.Android.CustomContent = ct
			r.Message.IOS = nil
		} else {
			checkIOSParams(r)
			r.Message.Android = nil
			r.Message.IOS.Custom = ct

		}
		return nil
	}
}

//OptionCustomContentSet 的注释
func OptionCustomContentSet(k, v string) PushMsgOption {
	return func(r *PushMsg) error {
		if r.Platform == PlatformAndroid {
			checkAndroidParams(r)
			if r.Message.Android.CustomContent == nil {
				r.Message.Android.CustomContent = map[string]string{k: v}
			} else {
				r.Message.Android.CustomContent[k] = v
			}
		} else {
			checkIOSParams(r)
			if r.Message.IOS.Custom == nil {
				r.Message.IOS.Custom = map[string]string{k: v}
			} else {
				r.Message.IOS.Custom[k] = v
			}
		}
		return nil
	}
}

//OptionMessage 的注释
func OptionMessage(m Message) PushMsgOption {
	return func(r *PushMsg) error {
		r.Message = m
		return nil
	}
}

//OptionTagList 的注释
func OptionTagList(op TagOperation, tags ...string) PushMsgOption {
	return func(r *PushMsg) error {
		r.TagList = &TagList{Tags: tags, TagOperation: op}
		return nil
	}
}

//OptionTagListOpt2 的注释
func OptionTagListOpt2(tl TagList) PushMsgOption {
	return func(r *PushMsg) error {
		r.TagList = &tl
		return nil
	}
}

//OptionTokenList 的注释
func OptionTokenList(tl ...string) PushMsgOption {
	return func(r *PushMsg) error {
		if len(tl) == 1 {
			r.AudienceType = AudienceTypeToken
		} else {
			r.AudienceType = AudienceTypeTokenList
		}
		r.TokenList = tl
		return nil
	}
}

//OptionTokenListAdd 的注释
func OptionTokenListAdd(t string) PushMsgOption {
	return func(r *PushMsg) error {
		if r.TokenList != nil {
			r.TokenList = append(r.TokenList, t)
		} else {
			r.TokenList = []string{t}
		}
		return nil
	}
}

//OptionAccountList 的注释
func OptionAccountList(al ...string) PushMsgOption {
	return func(r *PushMsg) error {
		if len(al) == 1 {
			r.AudienceType = AudienceTypeAccount
		} else {
			r.AudienceType = AudienceTypeAccountList
		}
		r.AccountList = al
		return nil
	}
}

//OptionAccountListAdd 的注释
func OptionAccountListAdd(a string) PushMsgOption {
	return func(r *PushMsg) error {
		if r.AccountList != nil {
			r.AccountList = append(r.AccountList, a)
		} else {
			r.AccountList = []string{a}
		}
		if len(r.AccountList) == 1 {
			r.AudienceType = AudienceTypeAccount
		} else {
			r.AudienceType = AudienceTypeAccountList
		}
		return nil
	}
}

//OptionExpireTime 的注释
func OptionExpireTime(et time.Time) PushMsgOption {
	return func(r *PushMsg) error {
		r.ExpireTime = int(et.Unix())
		return nil
	}
}

//OptionSendTime 修改发送时间
func OptionSendTime(st time.Time) PushMsgOption {
	return func(r *PushMsg) error {
		r.SendTime = st.Format("2006-01-02 15:04:05:07")
		return nil
	}
}

//OptionMultiPkg 的注释
func OptionMultiPkg(mp bool) PushMsgOption {
	return func(r *PushMsg) error {
		r.MultiPkg = mp
		return nil
	}
}

//OptionLoopTimes 的注释
func OptionLoopTimes(lt int) PushMsgOption {
	return func(r *PushMsg) error {
		r.LoopTimes = lt
		return nil
	}
}

//OptionStatTag 的注释
func OptionStatTag(st string) PushMsgOption {
	return func(r *PushMsg) error {
		r.StatTag = st
		return nil
	}
}

//OptionSeq 的注释
func OptionSeq(s int64) PushMsgOption {
	return func(r *PushMsg) error {
		r.Seq = s
		return nil
	}
}

//OptionAccountType 的注释
func OptionAccountType(at int) PushMsgOption {
	return func(r *PushMsg) error {
		r.AccountType = at
		return nil
	}
}

//OptionPushID 的注释
func OptionPushID(pid string) PushMsgOption {
	return func(r *PushMsg) error {
		r.PushID = pid
		return nil
	}
}

//OptionMessageType 的注释
func OptionMessageType(t MessageType) PushMsgOption {
	return func(r *PushMsg) error {
		r.MessageType = t
		return nil
	}
}

//OptionAps 的注释
func OptionAps(aps *Aps) PushMsgOption {
	return func(r *PushMsg) error {
		checkIOSParams(r)
		r.Message.IOS.Aps = aps
		return nil
	}
}

//OptionApsAlert 设置Alert
func OptionApsAlert(alert Alert) PushMsgOption {
	return func(r *PushMsg) error {
		checkIOSParams(r)
		r.Message.IOS.Aps.Alert = alert
		return nil
	}
}

//OptionApsBadage 设置Badage
func OptionApsBadage(badge int) PushMsgOption {
	return func(r *PushMsg) error {
		checkIOSParams(r)
		r.Message.IOS.Aps.Badge = badge
		return nil
	}
}

//OptionApsSound 设置属性
func OptionApsSound(sound string) PushMsgOption {
	return func(r *PushMsg) error {
		checkIOSParams(r)
		r.Message.IOS.Aps.Sound = sound
		return nil
	}
}

//OptionApsContentAvailable  设置属性
func OptionApsContentAvailable(contentAvailable int) PushMsgOption {
	return func(r *PushMsg) error {
		checkIOSParams(r)
		r.Message.IOS.Aps.ContentAvailable = contentAvailable
		return nil
	}
}

//OptionApsCategory  设置属性
func OptionApsCategory(category string) PushMsgOption {
	return func(r *PushMsg) error {
		checkIOSParams(r)
		r.Message.IOS.Aps.Category = category
		return nil
	}
}

//OptionApsThreadId  设置属性
func OptionApsThreadId(threadId string) PushMsgOption {
	return func(r *PushMsg) error {
		checkIOSParams(r)
		r.Message.IOS.Aps.ThreadId = threadId
		return nil
	}
}
