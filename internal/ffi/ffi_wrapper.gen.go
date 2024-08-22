/*------------------------------------------------------------------------------
//   This code was generated by template ffi_wrapper.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "ffi_wrapper.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/
package ffi

//revive:disable

// #include "gdextension_spx_codegen_header.h"
// #include "ffi_wrapper.gen.h"
// #include <stdint.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"

// C type aliases
// enums


// C function aliases
type GDExtensionSpxGlobalRegisterCallbacks C.GDExtensionSpxGlobalRegisterCallbacks
type GDExtensionSpxCallbackOnEngineStart C.GDExtensionSpxCallbackOnEngineStart
type GDExtensionSpxCallbackOnEngineUpdate C.GDExtensionSpxCallbackOnEngineUpdate
type GDExtensionSpxCallbackOnEngineDestroy C.GDExtensionSpxCallbackOnEngineDestroy
type GDExtensionSpxCallbackOnSpriteReady C.GDExtensionSpxCallbackOnSpriteReady
type GDExtensionSpxCallbackOnSpriteUpdated C.GDExtensionSpxCallbackOnSpriteUpdated
type GDExtensionSpxCallbackOnSpriteDestroyed C.GDExtensionSpxCallbackOnSpriteDestroyed
type GDExtensionSpxCallbackOnMousePressed C.GDExtensionSpxCallbackOnMousePressed
type GDExtensionSpxCallbackOnMouseReleased C.GDExtensionSpxCallbackOnMouseReleased
type GDExtensionSpxCallbackOnKeyPressed C.GDExtensionSpxCallbackOnKeyPressed
type GDExtensionSpxCallbackOnKeyReleased C.GDExtensionSpxCallbackOnKeyReleased
type GDExtensionSpxCallbackOnActionPressed C.GDExtensionSpxCallbackOnActionPressed
type GDExtensionSpxCallbackOnActionJustPressed C.GDExtensionSpxCallbackOnActionJustPressed
type GDExtensionSpxCallbackOnActionJustReleased C.GDExtensionSpxCallbackOnActionJustReleased
type GDExtensionSpxCallbackOnAxisChanged C.GDExtensionSpxCallbackOnAxisChanged
type GDExtensionSpxCallbackOnCollisionEnter C.GDExtensionSpxCallbackOnCollisionEnter
type GDExtensionSpxCallbackOnCollisionStay C.GDExtensionSpxCallbackOnCollisionStay
type GDExtensionSpxCallbackOnCollisionExit C.GDExtensionSpxCallbackOnCollisionExit
type GDExtensionSpxCallbackOnTriggerEnter C.GDExtensionSpxCallbackOnTriggerEnter
type GDExtensionSpxCallbackOnTriggerStay C.GDExtensionSpxCallbackOnTriggerStay
type GDExtensionSpxCallbackOnTriggerExit C.GDExtensionSpxCallbackOnTriggerExit
type GDExtensionSpxCallbackOnUIPressed C.GDExtensionSpxCallbackOnUIPressed
type GDExtensionSpxCallbackOnUIReleased C.GDExtensionSpxCallbackOnUIReleased
type GDExtensionSpxCallbackOnUIHovered C.GDExtensionSpxCallbackOnUIHovered
type GDExtensionSpxCallbackOnUIClicked C.GDExtensionSpxCallbackOnUIClicked
type GDExtensionSpxCallbackOnUIToggle C.GDExtensionSpxCallbackOnUIToggle
type GDExtensionSpxCallbackOnUITextChanged C.GDExtensionSpxCallbackOnUITextChanged
type GDExtensionSpxStringNewWithLatin1Chars C.GDExtensionSpxStringNewWithLatin1Chars
type GDExtensionSpxStringNewWithUtf8Chars C.GDExtensionSpxStringNewWithUtf8Chars
type GDExtensionSpxStringNewWithLatin1CharsAndLen C.GDExtensionSpxStringNewWithLatin1CharsAndLen
type GDExtensionSpxStringNewWithUtf8CharsAndLen C.GDExtensionSpxStringNewWithUtf8CharsAndLen
type GDExtensionSpxStringToLatin1Chars C.GDExtensionSpxStringToLatin1Chars
type GDExtensionSpxStringToUtf8Chars C.GDExtensionSpxStringToUtf8Chars
type GDExtensionSpxVariantGetPtrConstructor C.GDExtensionSpxVariantGetPtrConstructor
type GDExtensionSpxVariantGetPtrDestructor C.GDExtensionSpxVariantGetPtrDestructor
type GDExtensionSpxAudioPlayAudio C.GDExtensionSpxAudioPlayAudio
type GDExtensionSpxAudioSetAudioVolume C.GDExtensionSpxAudioSetAudioVolume
type GDExtensionSpxAudioGetAudioVolume C.GDExtensionSpxAudioGetAudioVolume
type GDExtensionSpxAudioIsMusicPlaying C.GDExtensionSpxAudioIsMusicPlaying
type GDExtensionSpxAudioPlayMusic C.GDExtensionSpxAudioPlayMusic
type GDExtensionSpxAudioSetMusicVolume C.GDExtensionSpxAudioSetMusicVolume
type GDExtensionSpxAudioGetMusicVolume C.GDExtensionSpxAudioGetMusicVolume
type GDExtensionSpxAudioPauseMusic C.GDExtensionSpxAudioPauseMusic
type GDExtensionSpxAudioResumeMusic C.GDExtensionSpxAudioResumeMusic
type GDExtensionSpxAudioGetMusicTimer C.GDExtensionSpxAudioGetMusicTimer
type GDExtensionSpxAudioSetMusicTimer C.GDExtensionSpxAudioSetMusicTimer
type GDExtensionSpxInputGetMousePos C.GDExtensionSpxInputGetMousePos
type GDExtensionSpxInputGetMouseState C.GDExtensionSpxInputGetMouseState
type GDExtensionSpxInputGetKeyState C.GDExtensionSpxInputGetKeyState
type GDExtensionSpxInputGetAxis C.GDExtensionSpxInputGetAxis
type GDExtensionSpxInputIsActionPressed C.GDExtensionSpxInputIsActionPressed
type GDExtensionSpxInputIsActionJustPressed C.GDExtensionSpxInputIsActionJustPressed
type GDExtensionSpxInputIsActionJustReleased C.GDExtensionSpxInputIsActionJustReleased
type GDExtensionSpxPhysicRaycast C.GDExtensionSpxPhysicRaycast
type GDExtensionSpxSpriteCreateSprite C.GDExtensionSpxSpriteCreateSprite
type GDExtensionSpxSpriteCloneSprite C.GDExtensionSpxSpriteCloneSprite
type GDExtensionSpxSpriteDestroySprite C.GDExtensionSpxSpriteDestroySprite
type GDExtensionSpxSpriteIsSpriteAlive C.GDExtensionSpxSpriteIsSpriteAlive
type GDExtensionSpxSpriteSetPosition C.GDExtensionSpxSpriteSetPosition
type GDExtensionSpxSpriteSetRotation C.GDExtensionSpxSpriteSetRotation
type GDExtensionSpxSpriteSetScale C.GDExtensionSpxSpriteSetScale
type GDExtensionSpxSpriteGetPosition C.GDExtensionSpxSpriteGetPosition
type GDExtensionSpxSpriteGetRotation C.GDExtensionSpxSpriteGetRotation
type GDExtensionSpxSpriteGetScale C.GDExtensionSpxSpriteGetScale
type GDExtensionSpxSpriteSetColor C.GDExtensionSpxSpriteSetColor
type GDExtensionSpxSpriteGetColor C.GDExtensionSpxSpriteGetColor
type GDExtensionSpxSpriteSetTexture C.GDExtensionSpxSpriteSetTexture
type GDExtensionSpxSpriteGetTexture C.GDExtensionSpxSpriteGetTexture
type GDExtensionSpxSpriteSetVisible C.GDExtensionSpxSpriteSetVisible
type GDExtensionSpxSpriteGetVisible C.GDExtensionSpxSpriteGetVisible
type GDExtensionSpxSpriteGetZIndex C.GDExtensionSpxSpriteGetZIndex
type GDExtensionSpxSpriteSetZIndex C.GDExtensionSpxSpriteSetZIndex
type GDExtensionSpxSpritePlayAnim C.GDExtensionSpxSpritePlayAnim
type GDExtensionSpxSpritePlayBackwardsAnim C.GDExtensionSpxSpritePlayBackwardsAnim
type GDExtensionSpxSpritePauseAnim C.GDExtensionSpxSpritePauseAnim
type GDExtensionSpxSpriteStopAnim C.GDExtensionSpxSpriteStopAnim
type GDExtensionSpxSpriteIsPlayingAnim C.GDExtensionSpxSpriteIsPlayingAnim
type GDExtensionSpxSpriteSetAnim C.GDExtensionSpxSpriteSetAnim
type GDExtensionSpxSpriteGetAnim C.GDExtensionSpxSpriteGetAnim
type GDExtensionSpxSpriteSetAnimFrame C.GDExtensionSpxSpriteSetAnimFrame
type GDExtensionSpxSpriteGetAnimFrame C.GDExtensionSpxSpriteGetAnimFrame
type GDExtensionSpxSpriteSetAnimSpeedScale C.GDExtensionSpxSpriteSetAnimSpeedScale
type GDExtensionSpxSpriteGetAnimSpeedScale C.GDExtensionSpxSpriteGetAnimSpeedScale
type GDExtensionSpxSpriteGetAnimPlayingSpeed C.GDExtensionSpxSpriteGetAnimPlayingSpeed
type GDExtensionSpxSpriteSetAnimCentered C.GDExtensionSpxSpriteSetAnimCentered
type GDExtensionSpxSpriteIsAnimCentered C.GDExtensionSpxSpriteIsAnimCentered
type GDExtensionSpxSpriteSetAnimOffset C.GDExtensionSpxSpriteSetAnimOffset
type GDExtensionSpxSpriteGetAnimOffset C.GDExtensionSpxSpriteGetAnimOffset
type GDExtensionSpxSpriteSetAnimFlipH C.GDExtensionSpxSpriteSetAnimFlipH
type GDExtensionSpxSpriteIsAnimFlippedH C.GDExtensionSpxSpriteIsAnimFlippedH
type GDExtensionSpxSpriteSetAnimFlipV C.GDExtensionSpxSpriteSetAnimFlipV
type GDExtensionSpxSpriteIsAnimFlippedV C.GDExtensionSpxSpriteIsAnimFlippedV
type GDExtensionSpxSpriteSetGravity C.GDExtensionSpxSpriteSetGravity
type GDExtensionSpxSpriteGetGravity C.GDExtensionSpxSpriteGetGravity
type GDExtensionSpxSpriteSetMass C.GDExtensionSpxSpriteSetMass
type GDExtensionSpxSpriteGetMass C.GDExtensionSpxSpriteGetMass
type GDExtensionSpxSpriteAddForce C.GDExtensionSpxSpriteAddForce
type GDExtensionSpxSpriteAddImpulse C.GDExtensionSpxSpriteAddImpulse
type GDExtensionSpxSpriteSetCollisionLayer C.GDExtensionSpxSpriteSetCollisionLayer
type GDExtensionSpxSpriteGetCollisionLayer C.GDExtensionSpxSpriteGetCollisionLayer
type GDExtensionSpxSpriteSetCollisionMask C.GDExtensionSpxSpriteSetCollisionMask
type GDExtensionSpxSpriteGetCollisionMask C.GDExtensionSpxSpriteGetCollisionMask
type GDExtensionSpxSpriteSetTriggerLayer C.GDExtensionSpxSpriteSetTriggerLayer
type GDExtensionSpxSpriteGetTriggerLayer C.GDExtensionSpxSpriteGetTriggerLayer
type GDExtensionSpxSpriteSetTriggerMask C.GDExtensionSpxSpriteSetTriggerMask
type GDExtensionSpxSpriteGetTriggerMask C.GDExtensionSpxSpriteGetTriggerMask
type GDExtensionSpxSpriteSetColliderRect C.GDExtensionSpxSpriteSetColliderRect
type GDExtensionSpxSpriteSetColliderCircle C.GDExtensionSpxSpriteSetColliderCircle
type GDExtensionSpxSpriteSetColliderCapsule C.GDExtensionSpxSpriteSetColliderCapsule
type GDExtensionSpxSpriteSetCollisionEnabled C.GDExtensionSpxSpriteSetCollisionEnabled
type GDExtensionSpxSpriteIsCollisionEnabled C.GDExtensionSpxSpriteIsCollisionEnabled
type GDExtensionSpxSpriteSetTriggerRect C.GDExtensionSpxSpriteSetTriggerRect
type GDExtensionSpxSpriteSetTriggerCircle C.GDExtensionSpxSpriteSetTriggerCircle
type GDExtensionSpxSpriteSetTriggerCapsule C.GDExtensionSpxSpriteSetTriggerCapsule
type GDExtensionSpxSpriteSetTriggerEnabled C.GDExtensionSpxSpriteSetTriggerEnabled
type GDExtensionSpxSpriteIsTriggerEnabled C.GDExtensionSpxSpriteIsTriggerEnabled
type GDExtensionSpxUICreateButton C.GDExtensionSpxUICreateButton
type GDExtensionSpxUICreateLabel C.GDExtensionSpxUICreateLabel
type GDExtensionSpxUICreateImage C.GDExtensionSpxUICreateImage
type GDExtensionSpxUICreateSlider C.GDExtensionSpxUICreateSlider
type GDExtensionSpxUICreateToggle C.GDExtensionSpxUICreateToggle
type GDExtensionSpxUICreateInput C.GDExtensionSpxUICreateInput
type GDExtensionSpxUIGetType C.GDExtensionSpxUIGetType
type GDExtensionSpxUISetInteractable C.GDExtensionSpxUISetInteractable
type GDExtensionSpxUIGetInteractable C.GDExtensionSpxUIGetInteractable
type GDExtensionSpxUISetText C.GDExtensionSpxUISetText
type GDExtensionSpxUIGetText C.GDExtensionSpxUIGetText
type GDExtensionSpxUISetRect C.GDExtensionSpxUISetRect
type GDExtensionSpxUIGetRect C.GDExtensionSpxUIGetRect
type GDExtensionSpxUISetColor C.GDExtensionSpxUISetColor
type GDExtensionSpxUIGetColor C.GDExtensionSpxUIGetColor
type GDExtensionSpxUISetFontSize C.GDExtensionSpxUISetFontSize
type GDExtensionSpxUIGetFontSize C.GDExtensionSpxUIGetFontSize
type GDExtensionSpxUISetVisible C.GDExtensionSpxUISetVisible
type GDExtensionSpxUIGetVisible C.GDExtensionSpxUIGetVisible


// call gdextension interface functions
func CallAudioPlayAudio(
	path GdString,
	)  {
	arg0 := (C.GDExtensionSpxAudioPlayAudio)(api.SpxAudioPlayAudio)
	arg1GdString = (C.GdString)(path)
	
	C.cgo_callfn_GDExtensionSpxAudioPlayAudio(arg0,arg1GdString,)
	
}
func CallAudioSetAudioVolume(
	volume GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxAudioSetAudioVolume)(api.SpxAudioSetAudioVolume)
	arg1GdFloat = (C.GdFloat)(volume)
	
	C.cgo_callfn_GDExtensionSpxAudioSetAudioVolume(arg0,arg1GdFloat,)
	
}
func CallAudioGetAudioVolume(
	) GdFloat {
	arg0 := (C.GDExtensionSpxAudioGetAudioVolume)(api.SpxAudioGetAudioVolume)
	var ret_val C.GdFloat 
	C.cgo_callfn_GDExtensionSpxAudioGetAudioVolume(arg0, &ret_val)
	return (GdFloat)(ret_val)
}
func CallAudioIsMusicPlaying(
	) GdBool {
	arg0 := (C.GDExtensionSpxAudioIsMusicPlaying)(api.SpxAudioIsMusicPlaying)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxAudioIsMusicPlaying(arg0, &ret_val)
	return (GdBool)(ret_val)
}
func CallAudioPlayMusic(
	path GdString,
	)  {
	arg0 := (C.GDExtensionSpxAudioPlayMusic)(api.SpxAudioPlayMusic)
	arg1GdString = (C.GdString)(path)
	
	C.cgo_callfn_GDExtensionSpxAudioPlayMusic(arg0,arg1GdString,)
	
}
func CallAudioSetMusicVolume(
	volume GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxAudioSetMusicVolume)(api.SpxAudioSetMusicVolume)
	arg1GdFloat = (C.GdFloat)(volume)
	
	C.cgo_callfn_GDExtensionSpxAudioSetMusicVolume(arg0,arg1GdFloat,)
	
}
func CallAudioGetMusicVolume(
	) GdFloat {
	arg0 := (C.GDExtensionSpxAudioGetMusicVolume)(api.SpxAudioGetMusicVolume)
	var ret_val C.GdFloat 
	C.cgo_callfn_GDExtensionSpxAudioGetMusicVolume(arg0, &ret_val)
	return (GdFloat)(ret_val)
}
func CallAudioPauseMusic(
	)  {
	arg0 := (C.GDExtensionSpxAudioPauseMusic)(api.SpxAudioPauseMusic)
	
	C.cgo_callfn_GDExtensionSpxAudioPauseMusic(arg0,)
}
func CallAudioResumeMusic(
	)  {
	arg0 := (C.GDExtensionSpxAudioResumeMusic)(api.SpxAudioResumeMusic)
	
	C.cgo_callfn_GDExtensionSpxAudioResumeMusic(arg0,)
}
func CallAudioGetMusicTimer(
	) GdFloat {
	arg0 := (C.GDExtensionSpxAudioGetMusicTimer)(api.SpxAudioGetMusicTimer)
	var ret_val C.GdFloat 
	C.cgo_callfn_GDExtensionSpxAudioGetMusicTimer(arg0, &ret_val)
	return (GdFloat)(ret_val)
}
func CallAudioSetMusicTimer(
	time GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxAudioSetMusicTimer)(api.SpxAudioSetMusicTimer)
	arg1GdFloat = (C.GdFloat)(time)
	
	C.cgo_callfn_GDExtensionSpxAudioSetMusicTimer(arg0,arg1GdFloat,)
	
}
func CallInputGetMousePos(
	) GdVec2 {
	arg0 := (C.GDExtensionSpxInputGetMousePos)(api.SpxInputGetMousePos)
	var ret_val C.GdVec2 
	C.cgo_callfn_GDExtensionSpxInputGetMousePos(arg0, &ret_val)
	return (GdVec2)(ret_val)
}
func CallInputGetMouseState(
	mouse_id GdInt,
	) GdBool {
	arg0 := (C.GDExtensionSpxInputGetMouseState)(api.SpxInputGetMouseState)
	arg1GdInt = (C.GdInt)(mouse_id)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxInputGetMouseState(arg0,arg1GdInt, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallInputGetKeyState(
	key GdInt,
	) GdInt {
	arg0 := (C.GDExtensionSpxInputGetKeyState)(api.SpxInputGetKeyState)
	arg1GdInt = (C.GdInt)(key)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxInputGetKeyState(arg0,arg1GdInt, &ret_val)
	
	return (GdInt)(ret_val)
}
func CallInputGetAxis(
	axis GdString,
	) GdFloat {
	arg0 := (C.GDExtensionSpxInputGetAxis)(api.SpxInputGetAxis)
	arg1GdString = (C.GdString)(axis)
	var ret_val C.GdFloat 
	C.cgo_callfn_GDExtensionSpxInputGetAxis(arg0,arg1GdString, &ret_val)
	
	return (GdFloat)(ret_val)
}
func CallInputIsActionPressed(
	action GdString,
	) GdBool {
	arg0 := (C.GDExtensionSpxInputIsActionPressed)(api.SpxInputIsActionPressed)
	arg1GdString = (C.GdString)(action)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxInputIsActionPressed(arg0,arg1GdString, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallInputIsActionJustPressed(
	action GdString,
	) GdBool {
	arg0 := (C.GDExtensionSpxInputIsActionJustPressed)(api.SpxInputIsActionJustPressed)
	arg1GdString = (C.GdString)(action)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxInputIsActionJustPressed(arg0,arg1GdString, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallInputIsActionJustReleased(
	action GdString,
	) GdBool {
	arg0 := (C.GDExtensionSpxInputIsActionJustReleased)(api.SpxInputIsActionJustReleased)
	arg1GdString = (C.GdString)(action)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxInputIsActionJustReleased(arg0,arg1GdString, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallPhysicRaycast(
	from GdVec2,
	to GdVec2,
	collision_mask GdInt,
	) GdObj {
	arg0 := (C.GDExtensionSpxPhysicRaycast)(api.SpxPhysicRaycast)
	arg1GdVec2 = (C.GdVec2)(from)
	arg2GdVec2 = (C.GdVec2)(to)
	arg3GdInt = (C.GdInt)(collision_mask)
	var ret_val C.GdObj 
	C.cgo_callfn_GDExtensionSpxPhysicRaycast(arg0,arg1GdVec2,arg2GdVec2,arg3GdInt, &ret_val)
	
	
	
	return (GdObj)(ret_val)
}
func CallSpriteCreateSprite(
	path GdString,
	) GdObj {
	arg0 := (C.GDExtensionSpxSpriteCreateSprite)(api.SpxSpriteCreateSprite)
	arg1GdString = (C.GdString)(path)
	var ret_val C.GdObj 
	C.cgo_callfn_GDExtensionSpxSpriteCreateSprite(arg0,arg1GdString, &ret_val)
	
	return (GdObj)(ret_val)
}
func CallSpriteCloneSprite(
	obj GdObj,
	) GdObj {
	arg0 := (C.GDExtensionSpxSpriteCloneSprite)(api.SpxSpriteCloneSprite)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdObj 
	C.cgo_callfn_GDExtensionSpxSpriteCloneSprite(arg0,arg1GdObj, &ret_val)
	
	return (GdObj)(ret_val)
}
func CallSpriteDestroySprite(
	obj GdObj,
	) GdBool {
	arg0 := (C.GDExtensionSpxSpriteDestroySprite)(api.SpxSpriteDestroySprite)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxSpriteDestroySprite(arg0,arg1GdObj, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallSpriteIsSpriteAlive(
	obj GdObj,
	) GdBool {
	arg0 := (C.GDExtensionSpxSpriteIsSpriteAlive)(api.SpxSpriteIsSpriteAlive)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxSpriteIsSpriteAlive(arg0,arg1GdObj, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallSpriteSetPosition(
	obj GdObj,
	pos GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetPosition)(api.SpxSpriteSetPosition)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdVec2 = (C.GdVec2)(pos)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetPosition(arg0,arg1GdObj,arg2GdVec2,)
	
	
}
func CallSpriteSetRotation(
	obj GdObj,
	rot GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetRotation)(api.SpxSpriteSetRotation)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdFloat = (C.GdFloat)(rot)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetRotation(arg0,arg1GdObj,arg2GdFloat,)
	
	
}
func CallSpriteSetScale(
	obj GdObj,
	scale GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetScale)(api.SpxSpriteSetScale)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdVec2 = (C.GdVec2)(scale)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetScale(arg0,arg1GdObj,arg2GdVec2,)
	
	
}
func CallSpriteGetPosition(
	obj GdObj,
	) GdVec2 {
	arg0 := (C.GDExtensionSpxSpriteGetPosition)(api.SpxSpriteGetPosition)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdVec2 
	C.cgo_callfn_GDExtensionSpxSpriteGetPosition(arg0,arg1GdObj, &ret_val)
	
	return (GdVec2)(ret_val)
}
func CallSpriteGetRotation(
	obj GdObj,
	) GdFloat {
	arg0 := (C.GDExtensionSpxSpriteGetRotation)(api.SpxSpriteGetRotation)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdFloat 
	C.cgo_callfn_GDExtensionSpxSpriteGetRotation(arg0,arg1GdObj, &ret_val)
	
	return (GdFloat)(ret_val)
}
func CallSpriteGetScale(
	obj GdObj,
	) GdVec2 {
	arg0 := (C.GDExtensionSpxSpriteGetScale)(api.SpxSpriteGetScale)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdVec2 
	C.cgo_callfn_GDExtensionSpxSpriteGetScale(arg0,arg1GdObj, &ret_val)
	
	return (GdVec2)(ret_val)
}
func CallSpriteSetColor(
	obj GdObj,
	color GdColor,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetColor)(api.SpxSpriteSetColor)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdColor = (C.GdColor)(color)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetColor(arg0,arg1GdObj,arg2GdColor,)
	
	
}
func CallSpriteGetColor(
	obj GdObj,
	) GdColor {
	arg0 := (C.GDExtensionSpxSpriteGetColor)(api.SpxSpriteGetColor)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdColor 
	C.cgo_callfn_GDExtensionSpxSpriteGetColor(arg0,arg1GdObj, &ret_val)
	
	return (GdColor)(ret_val)
}
func CallSpriteSetTexture(
	obj GdObj,
	path GdString,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetTexture)(api.SpxSpriteSetTexture)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdString = (C.GdString)(path)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetTexture(arg0,arg1GdObj,arg2GdString,)
	
	
}
func CallSpriteGetTexture(
	obj GdObj,
	) GdString {
	arg0 := (C.GDExtensionSpxSpriteGetTexture)(api.SpxSpriteGetTexture)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdString 
	C.cgo_callfn_GDExtensionSpxSpriteGetTexture(arg0,arg1GdObj, &ret_val)
	
	return (GdString)(ret_val)
}
func CallSpriteSetVisible(
	obj GdObj,
	visible GdBool,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetVisible)(api.SpxSpriteSetVisible)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdBool = (C.GdBool)(visible)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetVisible(arg0,arg1GdObj,arg2GdBool,)
	
	
}
func CallSpriteGetVisible(
	obj GdObj,
	) GdBool {
	arg0 := (C.GDExtensionSpxSpriteGetVisible)(api.SpxSpriteGetVisible)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxSpriteGetVisible(arg0,arg1GdObj, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallSpriteGetZIndex(
	obj GdObj,
	) GdInt {
	arg0 := (C.GDExtensionSpxSpriteGetZIndex)(api.SpxSpriteGetZIndex)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxSpriteGetZIndex(arg0,arg1GdObj, &ret_val)
	
	return (GdInt)(ret_val)
}
func CallSpriteSetZIndex(
	obj GdObj,
	z GdInt,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetZIndex)(api.SpxSpriteSetZIndex)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdInt = (C.GdInt)(z)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetZIndex(arg0,arg1GdObj,arg2GdInt,)
	
	
}
func CallSpritePlayAnim(
	obj GdObj,
	p_name GdString,
	p_custom_scale GdFloat,
	p_from_end GdBool,
	)  {
	arg0 := (C.GDExtensionSpxSpritePlayAnim)(api.SpxSpritePlayAnim)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdString = (C.GdString)(p_name)
	arg3GdFloat = (C.GdFloat)(p_custom_scale)
	arg4GdBool = (C.GdBool)(p_from_end)
	
	C.cgo_callfn_GDExtensionSpxSpritePlayAnim(arg0,arg1GdObj,arg2GdString,arg3GdFloat,arg4GdBool,)
	
	
	
	
}
func CallSpritePlayBackwardsAnim(
	obj GdObj,
	p_name GdString,
	)  {
	arg0 := (C.GDExtensionSpxSpritePlayBackwardsAnim)(api.SpxSpritePlayBackwardsAnim)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdString = (C.GdString)(p_name)
	
	C.cgo_callfn_GDExtensionSpxSpritePlayBackwardsAnim(arg0,arg1GdObj,arg2GdString,)
	
	
}
func CallSpritePauseAnim(
	obj GdObj,
	)  {
	arg0 := (C.GDExtensionSpxSpritePauseAnim)(api.SpxSpritePauseAnim)
	arg1GdObj = (C.GdObj)(obj)
	
	C.cgo_callfn_GDExtensionSpxSpritePauseAnim(arg0,arg1GdObj,)
	
}
func CallSpriteStopAnim(
	obj GdObj,
	)  {
	arg0 := (C.GDExtensionSpxSpriteStopAnim)(api.SpxSpriteStopAnim)
	arg1GdObj = (C.GdObj)(obj)
	
	C.cgo_callfn_GDExtensionSpxSpriteStopAnim(arg0,arg1GdObj,)
	
}
func CallSpriteIsPlayingAnim(
	obj GdObj,
	) GdBool {
	arg0 := (C.GDExtensionSpxSpriteIsPlayingAnim)(api.SpxSpriteIsPlayingAnim)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxSpriteIsPlayingAnim(arg0,arg1GdObj, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallSpriteSetAnim(
	obj GdObj,
	p_name GdString,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetAnim)(api.SpxSpriteSetAnim)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdString = (C.GdString)(p_name)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetAnim(arg0,arg1GdObj,arg2GdString,)
	
	
}
func CallSpriteGetAnim(
	obj GdObj,
	) GdString {
	arg0 := (C.GDExtensionSpxSpriteGetAnim)(api.SpxSpriteGetAnim)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdString 
	C.cgo_callfn_GDExtensionSpxSpriteGetAnim(arg0,arg1GdObj, &ret_val)
	
	return (GdString)(ret_val)
}
func CallSpriteSetAnimFrame(
	obj GdObj,
	p_frame GdInt,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetAnimFrame)(api.SpxSpriteSetAnimFrame)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdInt = (C.GdInt)(p_frame)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetAnimFrame(arg0,arg1GdObj,arg2GdInt,)
	
	
}
func CallSpriteGetAnimFrame(
	obj GdObj,
	) GdInt {
	arg0 := (C.GDExtensionSpxSpriteGetAnimFrame)(api.SpxSpriteGetAnimFrame)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxSpriteGetAnimFrame(arg0,arg1GdObj, &ret_val)
	
	return (GdInt)(ret_val)
}
func CallSpriteSetAnimSpeedScale(
	obj GdObj,
	p_speed_scale GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetAnimSpeedScale)(api.SpxSpriteSetAnimSpeedScale)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdFloat = (C.GdFloat)(p_speed_scale)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetAnimSpeedScale(arg0,arg1GdObj,arg2GdFloat,)
	
	
}
func CallSpriteGetAnimSpeedScale(
	obj GdObj,
	) GdFloat {
	arg0 := (C.GDExtensionSpxSpriteGetAnimSpeedScale)(api.SpxSpriteGetAnimSpeedScale)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdFloat 
	C.cgo_callfn_GDExtensionSpxSpriteGetAnimSpeedScale(arg0,arg1GdObj, &ret_val)
	
	return (GdFloat)(ret_val)
}
func CallSpriteGetAnimPlayingSpeed(
	obj GdObj,
	) GdFloat {
	arg0 := (C.GDExtensionSpxSpriteGetAnimPlayingSpeed)(api.SpxSpriteGetAnimPlayingSpeed)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdFloat 
	C.cgo_callfn_GDExtensionSpxSpriteGetAnimPlayingSpeed(arg0,arg1GdObj, &ret_val)
	
	return (GdFloat)(ret_val)
}
func CallSpriteSetAnimCentered(
	obj GdObj,
	p_center GdBool,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetAnimCentered)(api.SpxSpriteSetAnimCentered)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdBool = (C.GdBool)(p_center)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetAnimCentered(arg0,arg1GdObj,arg2GdBool,)
	
	
}
func CallSpriteIsAnimCentered(
	obj GdObj,
	) GdBool {
	arg0 := (C.GDExtensionSpxSpriteIsAnimCentered)(api.SpxSpriteIsAnimCentered)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxSpriteIsAnimCentered(arg0,arg1GdObj, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallSpriteSetAnimOffset(
	obj GdObj,
	p_offset GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetAnimOffset)(api.SpxSpriteSetAnimOffset)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdVec2 = (C.GdVec2)(p_offset)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetAnimOffset(arg0,arg1GdObj,arg2GdVec2,)
	
	
}
func CallSpriteGetAnimOffset(
	obj GdObj,
	) GdVec2 {
	arg0 := (C.GDExtensionSpxSpriteGetAnimOffset)(api.SpxSpriteGetAnimOffset)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdVec2 
	C.cgo_callfn_GDExtensionSpxSpriteGetAnimOffset(arg0,arg1GdObj, &ret_val)
	
	return (GdVec2)(ret_val)
}
func CallSpriteSetAnimFlipH(
	obj GdObj,
	p_flip GdBool,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetAnimFlipH)(api.SpxSpriteSetAnimFlipH)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdBool = (C.GdBool)(p_flip)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetAnimFlipH(arg0,arg1GdObj,arg2GdBool,)
	
	
}
func CallSpriteIsAnimFlippedH(
	obj GdObj,
	) GdBool {
	arg0 := (C.GDExtensionSpxSpriteIsAnimFlippedH)(api.SpxSpriteIsAnimFlippedH)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxSpriteIsAnimFlippedH(arg0,arg1GdObj, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallSpriteSetAnimFlipV(
	obj GdObj,
	p_flip GdBool,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetAnimFlipV)(api.SpxSpriteSetAnimFlipV)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdBool = (C.GdBool)(p_flip)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetAnimFlipV(arg0,arg1GdObj,arg2GdBool,)
	
	
}
func CallSpriteIsAnimFlippedV(
	obj GdObj,
	) GdBool {
	arg0 := (C.GDExtensionSpxSpriteIsAnimFlippedV)(api.SpxSpriteIsAnimFlippedV)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxSpriteIsAnimFlippedV(arg0,arg1GdObj, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallSpriteSetGravity(
	obj GdObj,
	gravity GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetGravity)(api.SpxSpriteSetGravity)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdFloat = (C.GdFloat)(gravity)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetGravity(arg0,arg1GdObj,arg2GdFloat,)
	
	
}
func CallSpriteGetGravity(
	obj GdObj,
	) GdFloat {
	arg0 := (C.GDExtensionSpxSpriteGetGravity)(api.SpxSpriteGetGravity)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdFloat 
	C.cgo_callfn_GDExtensionSpxSpriteGetGravity(arg0,arg1GdObj, &ret_val)
	
	return (GdFloat)(ret_val)
}
func CallSpriteSetMass(
	obj GdObj,
	mass GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetMass)(api.SpxSpriteSetMass)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdFloat = (C.GdFloat)(mass)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetMass(arg0,arg1GdObj,arg2GdFloat,)
	
	
}
func CallSpriteGetMass(
	obj GdObj,
	) GdFloat {
	arg0 := (C.GDExtensionSpxSpriteGetMass)(api.SpxSpriteGetMass)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdFloat 
	C.cgo_callfn_GDExtensionSpxSpriteGetMass(arg0,arg1GdObj, &ret_val)
	
	return (GdFloat)(ret_val)
}
func CallSpriteAddForce(
	obj GdObj,
	force GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxSpriteAddForce)(api.SpxSpriteAddForce)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdVec2 = (C.GdVec2)(force)
	
	C.cgo_callfn_GDExtensionSpxSpriteAddForce(arg0,arg1GdObj,arg2GdVec2,)
	
	
}
func CallSpriteAddImpulse(
	obj GdObj,
	impulse GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxSpriteAddImpulse)(api.SpxSpriteAddImpulse)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdVec2 = (C.GdVec2)(impulse)
	
	C.cgo_callfn_GDExtensionSpxSpriteAddImpulse(arg0,arg1GdObj,arg2GdVec2,)
	
	
}
func CallSpriteSetCollisionLayer(
	obj GdObj,
	layer GdInt,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetCollisionLayer)(api.SpxSpriteSetCollisionLayer)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdInt = (C.GdInt)(layer)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetCollisionLayer(arg0,arg1GdObj,arg2GdInt,)
	
	
}
func CallSpriteGetCollisionLayer(
	obj GdObj,
	) GdInt {
	arg0 := (C.GDExtensionSpxSpriteGetCollisionLayer)(api.SpxSpriteGetCollisionLayer)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxSpriteGetCollisionLayer(arg0,arg1GdObj, &ret_val)
	
	return (GdInt)(ret_val)
}
func CallSpriteSetCollisionMask(
	obj GdObj,
	mask GdInt,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetCollisionMask)(api.SpxSpriteSetCollisionMask)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdInt = (C.GdInt)(mask)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetCollisionMask(arg0,arg1GdObj,arg2GdInt,)
	
	
}
func CallSpriteGetCollisionMask(
	obj GdObj,
	) GdInt {
	arg0 := (C.GDExtensionSpxSpriteGetCollisionMask)(api.SpxSpriteGetCollisionMask)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxSpriteGetCollisionMask(arg0,arg1GdObj, &ret_val)
	
	return (GdInt)(ret_val)
}
func CallSpriteSetTriggerLayer(
	obj GdObj,
	layer GdInt,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetTriggerLayer)(api.SpxSpriteSetTriggerLayer)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdInt = (C.GdInt)(layer)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetTriggerLayer(arg0,arg1GdObj,arg2GdInt,)
	
	
}
func CallSpriteGetTriggerLayer(
	obj GdObj,
	) GdInt {
	arg0 := (C.GDExtensionSpxSpriteGetTriggerLayer)(api.SpxSpriteGetTriggerLayer)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxSpriteGetTriggerLayer(arg0,arg1GdObj, &ret_val)
	
	return (GdInt)(ret_val)
}
func CallSpriteSetTriggerMask(
	obj GdObj,
	mask GdInt,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetTriggerMask)(api.SpxSpriteSetTriggerMask)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdInt = (C.GdInt)(mask)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetTriggerMask(arg0,arg1GdObj,arg2GdInt,)
	
	
}
func CallSpriteGetTriggerMask(
	obj GdObj,
	) GdInt {
	arg0 := (C.GDExtensionSpxSpriteGetTriggerMask)(api.SpxSpriteGetTriggerMask)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxSpriteGetTriggerMask(arg0,arg1GdObj, &ret_val)
	
	return (GdInt)(ret_val)
}
func CallSpriteSetColliderRect(
	obj GdObj,
	center GdVec2,
	size GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetColliderRect)(api.SpxSpriteSetColliderRect)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdVec2 = (C.GdVec2)(center)
	arg3GdVec2 = (C.GdVec2)(size)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetColliderRect(arg0,arg1GdObj,arg2GdVec2,arg3GdVec2,)
	
	
	
}
func CallSpriteSetColliderCircle(
	obj GdObj,
	center GdVec2,
	radius GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetColliderCircle)(api.SpxSpriteSetColliderCircle)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdVec2 = (C.GdVec2)(center)
	arg3GdFloat = (C.GdFloat)(radius)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetColliderCircle(arg0,arg1GdObj,arg2GdVec2,arg3GdFloat,)
	
	
	
}
func CallSpriteSetColliderCapsule(
	obj GdObj,
	center GdVec2,
	size GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetColliderCapsule)(api.SpxSpriteSetColliderCapsule)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdVec2 = (C.GdVec2)(center)
	arg3GdVec2 = (C.GdVec2)(size)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetColliderCapsule(arg0,arg1GdObj,arg2GdVec2,arg3GdVec2,)
	
	
	
}
func CallSpriteSetCollisionEnabled(
	obj GdObj,
	enabled GdBool,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetCollisionEnabled)(api.SpxSpriteSetCollisionEnabled)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdBool = (C.GdBool)(enabled)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetCollisionEnabled(arg0,arg1GdObj,arg2GdBool,)
	
	
}
func CallSpriteIsCollisionEnabled(
	obj GdObj,
	) GdBool {
	arg0 := (C.GDExtensionSpxSpriteIsCollisionEnabled)(api.SpxSpriteIsCollisionEnabled)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxSpriteIsCollisionEnabled(arg0,arg1GdObj, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallSpriteSetTriggerRect(
	obj GdObj,
	center GdVec2,
	size GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetTriggerRect)(api.SpxSpriteSetTriggerRect)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdVec2 = (C.GdVec2)(center)
	arg3GdVec2 = (C.GdVec2)(size)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetTriggerRect(arg0,arg1GdObj,arg2GdVec2,arg3GdVec2,)
	
	
	
}
func CallSpriteSetTriggerCircle(
	obj GdObj,
	center GdVec2,
	radius GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetTriggerCircle)(api.SpxSpriteSetTriggerCircle)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdVec2 = (C.GdVec2)(center)
	arg3GdFloat = (C.GdFloat)(radius)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetTriggerCircle(arg0,arg1GdObj,arg2GdVec2,arg3GdFloat,)
	
	
	
}
func CallSpriteSetTriggerCapsule(
	obj GdObj,
	center GdVec2,
	size GdVec2,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetTriggerCapsule)(api.SpxSpriteSetTriggerCapsule)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdVec2 = (C.GdVec2)(center)
	arg3GdVec2 = (C.GdVec2)(size)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetTriggerCapsule(arg0,arg1GdObj,arg2GdVec2,arg3GdVec2,)
	
	
	
}
func CallSpriteSetTriggerEnabled(
	obj GdObj,
	trigger GdBool,
	)  {
	arg0 := (C.GDExtensionSpxSpriteSetTriggerEnabled)(api.SpxSpriteSetTriggerEnabled)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdBool = (C.GdBool)(trigger)
	
	C.cgo_callfn_GDExtensionSpxSpriteSetTriggerEnabled(arg0,arg1GdObj,arg2GdBool,)
	
	
}
func CallSpriteIsTriggerEnabled(
	obj GdObj,
	) GdBool {
	arg0 := (C.GDExtensionSpxSpriteIsTriggerEnabled)(api.SpxSpriteIsTriggerEnabled)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxSpriteIsTriggerEnabled(arg0,arg1GdObj, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallUICreateButton(
	path GdString,
	rect GdRect2,
	text GdString,
	) GdInt {
	arg0 := (C.GDExtensionSpxUICreateButton)(api.SpxUICreateButton)
	arg1GdString = (C.GdString)(path)
	arg2GdRect2 = (C.GdRect2)(rect)
	arg3GdString = (C.GdString)(text)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxUICreateButton(arg0,arg1GdString,arg2GdRect2,arg3GdString, &ret_val)
	
	
	
	return (GdInt)(ret_val)
}
func CallUICreateLabel(
	path GdString,
	rect GdRect2,
	text GdString,
	) GdInt {
	arg0 := (C.GDExtensionSpxUICreateLabel)(api.SpxUICreateLabel)
	arg1GdString = (C.GdString)(path)
	arg2GdRect2 = (C.GdRect2)(rect)
	arg3GdString = (C.GdString)(text)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxUICreateLabel(arg0,arg1GdString,arg2GdRect2,arg3GdString, &ret_val)
	
	
	
	return (GdInt)(ret_val)
}
func CallUICreateImage(
	path GdString,
	rect GdRect2,
	color GdColor,
	) GdInt {
	arg0 := (C.GDExtensionSpxUICreateImage)(api.SpxUICreateImage)
	arg1GdString = (C.GdString)(path)
	arg2GdRect2 = (C.GdRect2)(rect)
	arg3GdColor = (C.GdColor)(color)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxUICreateImage(arg0,arg1GdString,arg2GdRect2,arg3GdColor, &ret_val)
	
	
	
	return (GdInt)(ret_val)
}
func CallUICreateSlider(
	path GdString,
	rect GdRect2,
	value GdFloat,
	) GdInt {
	arg0 := (C.GDExtensionSpxUICreateSlider)(api.SpxUICreateSlider)
	arg1GdString = (C.GdString)(path)
	arg2GdRect2 = (C.GdRect2)(rect)
	arg3GdFloat = (C.GdFloat)(value)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxUICreateSlider(arg0,arg1GdString,arg2GdRect2,arg3GdFloat, &ret_val)
	
	
	
	return (GdInt)(ret_val)
}
func CallUICreateToggle(
	path GdString,
	rect GdRect2,
	value GdBool,
	) GdInt {
	arg0 := (C.GDExtensionSpxUICreateToggle)(api.SpxUICreateToggle)
	arg1GdString = (C.GdString)(path)
	arg2GdRect2 = (C.GdRect2)(rect)
	arg3GdBool = (C.GdBool)(value)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxUICreateToggle(arg0,arg1GdString,arg2GdRect2,arg3GdBool, &ret_val)
	
	
	
	return (GdInt)(ret_val)
}
func CallUICreateInput(
	path GdString,
	rect GdRect2,
	text GdString,
	) GdInt {
	arg0 := (C.GDExtensionSpxUICreateInput)(api.SpxUICreateInput)
	arg1GdString = (C.GdString)(path)
	arg2GdRect2 = (C.GdRect2)(rect)
	arg3GdString = (C.GdString)(text)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxUICreateInput(arg0,arg1GdString,arg2GdRect2,arg3GdString, &ret_val)
	
	
	
	return (GdInt)(ret_val)
}
func CallUIGetType(
	obj GdObj,
	) GdInt {
	arg0 := (C.GDExtensionSpxUIGetType)(api.SpxUIGetType)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdInt 
	C.cgo_callfn_GDExtensionSpxUIGetType(arg0,arg1GdObj, &ret_val)
	
	return (GdInt)(ret_val)
}
func CallUISetInteractable(
	obj GdObj,
	interactable GdBool,
	)  {
	arg0 := (C.GDExtensionSpxUISetInteractable)(api.SpxUISetInteractable)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdBool = (C.GdBool)(interactable)
	
	C.cgo_callfn_GDExtensionSpxUISetInteractable(arg0,arg1GdObj,arg2GdBool,)
	
	
}
func CallUIGetInteractable(
	obj GdObj,
	) GdBool {
	arg0 := (C.GDExtensionSpxUIGetInteractable)(api.SpxUIGetInteractable)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxUIGetInteractable(arg0,arg1GdObj, &ret_val)
	
	return (GdBool)(ret_val)
}
func CallUISetText(
	obj GdObj,
	text GdString,
	)  {
	arg0 := (C.GDExtensionSpxUISetText)(api.SpxUISetText)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdString = (C.GdString)(text)
	
	C.cgo_callfn_GDExtensionSpxUISetText(arg0,arg1GdObj,arg2GdString,)
	
	
}
func CallUIGetText(
	obj GdObj,
	) GdString {
	arg0 := (C.GDExtensionSpxUIGetText)(api.SpxUIGetText)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdString 
	C.cgo_callfn_GDExtensionSpxUIGetText(arg0,arg1GdObj, &ret_val)
	
	return (GdString)(ret_val)
}
func CallUISetRect(
	obj GdObj,
	rect GdRect2,
	)  {
	arg0 := (C.GDExtensionSpxUISetRect)(api.SpxUISetRect)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdRect2 = (C.GdRect2)(rect)
	
	C.cgo_callfn_GDExtensionSpxUISetRect(arg0,arg1GdObj,arg2GdRect2,)
	
	
}
func CallUIGetRect(
	obj GdObj,
	) GdRect2 {
	arg0 := (C.GDExtensionSpxUIGetRect)(api.SpxUIGetRect)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdRect2 
	C.cgo_callfn_GDExtensionSpxUIGetRect(arg0,arg1GdObj, &ret_val)
	
	return (GdRect2)(ret_val)
}
func CallUISetColor(
	obj GdObj,
	color GdColor,
	)  {
	arg0 := (C.GDExtensionSpxUISetColor)(api.SpxUISetColor)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdColor = (C.GdColor)(color)
	
	C.cgo_callfn_GDExtensionSpxUISetColor(arg0,arg1GdObj,arg2GdColor,)
	
	
}
func CallUIGetColor(
	obj GdObj,
	) GdColor {
	arg0 := (C.GDExtensionSpxUIGetColor)(api.SpxUIGetColor)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdColor 
	C.cgo_callfn_GDExtensionSpxUIGetColor(arg0,arg1GdObj, &ret_val)
	
	return (GdColor)(ret_val)
}
func CallUISetFontSize(
	obj GdObj,
	size GdFloat,
	)  {
	arg0 := (C.GDExtensionSpxUISetFontSize)(api.SpxUISetFontSize)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdFloat = (C.GdFloat)(size)
	
	C.cgo_callfn_GDExtensionSpxUISetFontSize(arg0,arg1GdObj,arg2GdFloat,)
	
	
}
func CallUIGetFontSize(
	obj GdObj,
	) GdFloat {
	arg0 := (C.GDExtensionSpxUIGetFontSize)(api.SpxUIGetFontSize)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdFloat 
	C.cgo_callfn_GDExtensionSpxUIGetFontSize(arg0,arg1GdObj, &ret_val)
	
	return (GdFloat)(ret_val)
}
func CallUISetVisible(
	obj GdObj,
	visible GdBool,
	)  {
	arg0 := (C.GDExtensionSpxUISetVisible)(api.SpxUISetVisible)
	arg1GdObj = (C.GdObj)(obj)
	arg2GdBool = (C.GdBool)(visible)
	
	C.cgo_callfn_GDExtensionSpxUISetVisible(arg0,arg1GdObj,arg2GdBool,)
	
	
}
func CallUIGetVisible(
	obj GdObj,
	) GdBool {
	arg0 := (C.GDExtensionSpxUIGetVisible)(api.SpxUIGetVisible)
	arg1GdObj = (C.GdObj)(obj)
	var ret_val C.GdBool 
	C.cgo_callfn_GDExtensionSpxUIGetVisible(arg0,arg1GdObj, &ret_val)
	
	return (GdBool)(ret_val)
}