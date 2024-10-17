#ifndef CGO_GODOT_GO_GDEXTENSION_WRAPPER_H
#define CGO_GODOT_GO_GDEXTENSION_WRAPPER_H


/*------------------------------------------------------------------------------
//   This code was generated by template ffi_wrapper.c.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "ffi_wrapper.h.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/
#include "gdextension_spx_codegen_header.h"

/* Go cannot call C function pointers directly, so we must generate C wrapper code to call the functions. */void cgo_callfn_GDExtensionSpxAudioStopAll(const GDExtensionSpxAudioStopAll fn) {
	fn();
}
void cgo_callfn_GDExtensionSpxAudioPlaySfx(const GDExtensionSpxAudioPlaySfx fn, GdString path) {
	fn(path);
}
void cgo_callfn_GDExtensionSpxAudioPlayMusic(const GDExtensionSpxAudioPlayMusic fn, GdString path) {
	fn(path);
}
void cgo_callfn_GDExtensionSpxAudioPauseMusic(const GDExtensionSpxAudioPauseMusic fn) {
	fn();
}
void cgo_callfn_GDExtensionSpxAudioResumeMusic(const GDExtensionSpxAudioResumeMusic fn) {
	fn();
}
void cgo_callfn_GDExtensionSpxAudioGetMusicTimer(const GDExtensionSpxAudioGetMusicTimer fn, GdFloat* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxAudioSetMusicTimer(const GDExtensionSpxAudioSetMusicTimer fn, GdFloat time) {
	fn(time);
}
void cgo_callfn_GDExtensionSpxAudioIsMusicPlaying(const GDExtensionSpxAudioIsMusicPlaying fn, GdBool* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxAudioSetSfxVolume(const GDExtensionSpxAudioSetSfxVolume fn, GdFloat volume) {
	fn(volume);
}
void cgo_callfn_GDExtensionSpxAudioGetSfxVolume(const GDExtensionSpxAudioGetSfxVolume fn, GdFloat* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxAudioSetMusicVolume(const GDExtensionSpxAudioSetMusicVolume fn, GdFloat volume) {
	fn(volume);
}
void cgo_callfn_GDExtensionSpxAudioGetMusicVolume(const GDExtensionSpxAudioGetMusicVolume fn, GdFloat* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxAudioSetMasterVolume(const GDExtensionSpxAudioSetMasterVolume fn, GdFloat volume) {
	fn(volume);
}
void cgo_callfn_GDExtensionSpxAudioGetMasterVolume(const GDExtensionSpxAudioGetMasterVolume fn, GdFloat* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxCameraGetCameraPosition(const GDExtensionSpxCameraGetCameraPosition fn, GdVec2* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxCameraSetCameraPosition(const GDExtensionSpxCameraSetCameraPosition fn, GdVec2 position) {
	fn(position);
}
void cgo_callfn_GDExtensionSpxCameraGetCameraZoom(const GDExtensionSpxCameraGetCameraZoom fn, GdVec2* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxCameraSetCameraZoom(const GDExtensionSpxCameraSetCameraZoom fn, GdVec2 size) {
	fn(size);
}
void cgo_callfn_GDExtensionSpxCameraGetViewportRect(const GDExtensionSpxCameraGetViewportRect fn, GdRect2* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxInputGetMousePos(const GDExtensionSpxInputGetMousePos fn, GdVec2* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxInputGetKey(const GDExtensionSpxInputGetKey fn, GdInt key, GdBool* ret_val) {
	fn(key,ret_val);
}
void cgo_callfn_GDExtensionSpxInputGetMouseState(const GDExtensionSpxInputGetMouseState fn, GdInt mouse_id, GdBool* ret_val) {
	fn(mouse_id,ret_val);
}
void cgo_callfn_GDExtensionSpxInputGetKeyState(const GDExtensionSpxInputGetKeyState fn, GdInt key, GdInt* ret_val) {
	fn(key,ret_val);
}
void cgo_callfn_GDExtensionSpxInputGetAxis(const GDExtensionSpxInputGetAxis fn, GdString neg_action, GdString pos_action, GdFloat* ret_val) {
	fn(neg_action, pos_action,ret_val);
}
void cgo_callfn_GDExtensionSpxInputIsActionPressed(const GDExtensionSpxInputIsActionPressed fn, GdString action, GdBool* ret_val) {
	fn(action,ret_val);
}
void cgo_callfn_GDExtensionSpxInputIsActionJustPressed(const GDExtensionSpxInputIsActionJustPressed fn, GdString action, GdBool* ret_val) {
	fn(action,ret_val);
}
void cgo_callfn_GDExtensionSpxInputIsActionJustReleased(const GDExtensionSpxInputIsActionJustReleased fn, GdString action, GdBool* ret_val) {
	fn(action,ret_val);
}
void cgo_callfn_GDExtensionSpxPhysicRaycast(const GDExtensionSpxPhysicRaycast fn, GdVec2 from, GdVec2 to, GdInt collision_mask, GdObj* ret_val) {
	fn(from, to, collision_mask,ret_val);
}
void cgo_callfn_GDExtensionSpxPhysicCheckCollision(const GDExtensionSpxPhysicCheckCollision fn, GdVec2 from, GdVec2 to, GdInt collision_mask, GdBool collide_with_areas, GdBool collide_with_bodies, GdBool* ret_val) {
	fn(from, to, collision_mask, collide_with_areas, collide_with_bodies,ret_val);
}
void cgo_callfn_GDExtensionSpxPlatformSetWindowSize(const GDExtensionSpxPlatformSetWindowSize fn, GdInt width, GdInt height) {
	fn(width, height);
}
void cgo_callfn_GDExtensionSpxPlatformGetWindowSize(const GDExtensionSpxPlatformGetWindowSize fn, GdVec2* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxPlatformSetWindowTitle(const GDExtensionSpxPlatformSetWindowTitle fn, GdString title) {
	fn(title);
}
void cgo_callfn_GDExtensionSpxPlatformGetWindowTitle(const GDExtensionSpxPlatformGetWindowTitle fn, GdString* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxPlatformSetWindowFullscreen(const GDExtensionSpxPlatformSetWindowFullscreen fn, GdBool enable) {
	fn(enable);
}
void cgo_callfn_GDExtensionSpxPlatformIsWindowFullscreen(const GDExtensionSpxPlatformIsWindowFullscreen fn, GdBool* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxPlatformSetDebugMode(const GDExtensionSpxPlatformSetDebugMode fn, GdBool enable) {
	fn(enable);
}
void cgo_callfn_GDExtensionSpxPlatformIsDebugMode(const GDExtensionSpxPlatformIsDebugMode fn, GdBool* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxResGetImageSize(const GDExtensionSpxResGetImageSize fn, GdString path, GdVec2* ret_val) {
	fn(path,ret_val);
}
void cgo_callfn_GDExtensionSpxResReadAllText(const GDExtensionSpxResReadAllText fn, GdString path, GdString* ret_val) {
	fn(path,ret_val);
}
void cgo_callfn_GDExtensionSpxSceneChangeSceneToFile(const GDExtensionSpxSceneChangeSceneToFile fn, GdString path) {
	fn(path);
}
void cgo_callfn_GDExtensionSpxSceneReloadCurrentScene(const GDExtensionSpxSceneReloadCurrentScene fn, GdInt* ret_val) {
	fn(ret_val);
}
void cgo_callfn_GDExtensionSpxSceneUnloadCurrentScene(const GDExtensionSpxSceneUnloadCurrentScene fn) {
	fn();
}
void cgo_callfn_GDExtensionSpxSpriteSetDontDestroyOnLoad(const GDExtensionSpxSpriteSetDontDestroyOnLoad fn, GdObj obj) {
	fn(obj);
}
void cgo_callfn_GDExtensionSpxSpriteSetProcess(const GDExtensionSpxSpriteSetProcess fn, GdObj obj, GdBool is_on) {
	fn(obj, is_on);
}
void cgo_callfn_GDExtensionSpxSpriteSetPhysicProcess(const GDExtensionSpxSpriteSetPhysicProcess fn, GdObj obj, GdBool is_on) {
	fn(obj, is_on);
}
void cgo_callfn_GDExtensionSpxSpriteSetChildPosition(const GDExtensionSpxSpriteSetChildPosition fn, GdObj obj, GdString path, GdVec2 pos) {
	fn(obj, path, pos);
}
void cgo_callfn_GDExtensionSpxSpriteGetChildPosition(const GDExtensionSpxSpriteGetChildPosition fn, GdObj obj, GdString path, GdVec2* ret_val) {
	fn(obj, path,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetChildRotation(const GDExtensionSpxSpriteSetChildRotation fn, GdObj obj, GdString path, GdFloat rot) {
	fn(obj, path, rot);
}
void cgo_callfn_GDExtensionSpxSpriteGetChildRotation(const GDExtensionSpxSpriteGetChildRotation fn, GdObj obj, GdString path, GdFloat* ret_val) {
	fn(obj, path,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetChildScale(const GDExtensionSpxSpriteSetChildScale fn, GdObj obj, GdString path, GdVec2 scale) {
	fn(obj, path, scale);
}
void cgo_callfn_GDExtensionSpxSpriteGetChildScale(const GDExtensionSpxSpriteGetChildScale fn, GdObj obj, GdString path, GdVec2* ret_val) {
	fn(obj, path,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteCheckCollision(const GDExtensionSpxSpriteCheckCollision fn, GdObj obj, GdObj target, GdBool is_src_trigger, GdBool is_dst_trigger, GdBool* ret_val) {
	fn(obj, target, is_src_trigger, is_dst_trigger,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteCheckCollisionWithPoint(const GDExtensionSpxSpriteCheckCollisionWithPoint fn, GdObj obj, GdVec2 point, GdBool is_trigger, GdBool* ret_val) {
	fn(obj, point, is_trigger,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteCreateSprite(const GDExtensionSpxSpriteCreateSprite fn, GdString path, GdObj* ret_val) {
	fn(path,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteCloneSprite(const GDExtensionSpxSpriteCloneSprite fn, GdObj obj, GdObj* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteDestroySprite(const GDExtensionSpxSpriteDestroySprite fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteIsSpriteAlive(const GDExtensionSpxSpriteIsSpriteAlive fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetPosition(const GDExtensionSpxSpriteSetPosition fn, GdObj obj, GdVec2 pos) {
	fn(obj, pos);
}
void cgo_callfn_GDExtensionSpxSpriteGetPosition(const GDExtensionSpxSpriteGetPosition fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetRotation(const GDExtensionSpxSpriteSetRotation fn, GdObj obj, GdFloat rot) {
	fn(obj, rot);
}
void cgo_callfn_GDExtensionSpxSpriteGetRotation(const GDExtensionSpxSpriteGetRotation fn, GdObj obj, GdFloat* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetScale(const GDExtensionSpxSpriteSetScale fn, GdObj obj, GdVec2 scale) {
	fn(obj, scale);
}
void cgo_callfn_GDExtensionSpxSpriteGetScale(const GDExtensionSpxSpriteGetScale fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetRenderScale(const GDExtensionSpxSpriteSetRenderScale fn, GdObj obj, GdVec2 scale) {
	fn(obj, scale);
}
void cgo_callfn_GDExtensionSpxSpriteGetRenderScale(const GDExtensionSpxSpriteGetRenderScale fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetColor(const GDExtensionSpxSpriteSetColor fn, GdObj obj, GdColor color) {
	fn(obj, color);
}
void cgo_callfn_GDExtensionSpxSpriteGetColor(const GDExtensionSpxSpriteGetColor fn, GdObj obj, GdColor* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetTextureAltas(const GDExtensionSpxSpriteSetTextureAltas fn, GdObj obj, GdString path, GdRect2 rect2) {
	fn(obj, path, rect2);
}
void cgo_callfn_GDExtensionSpxSpriteSetTexture(const GDExtensionSpxSpriteSetTexture fn, GdObj obj, GdString path) {
	fn(obj, path);
}
void cgo_callfn_GDExtensionSpxSpriteGetTexture(const GDExtensionSpxSpriteGetTexture fn, GdObj obj, GdString* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetVisible(const GDExtensionSpxSpriteSetVisible fn, GdObj obj, GdBool visible) {
	fn(obj, visible);
}
void cgo_callfn_GDExtensionSpxSpriteGetVisible(const GDExtensionSpxSpriteGetVisible fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteGetZIndex(const GDExtensionSpxSpriteGetZIndex fn, GdObj obj, GdInt* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetZIndex(const GDExtensionSpxSpriteSetZIndex fn, GdObj obj, GdInt z) {
	fn(obj, z);
}
void cgo_callfn_GDExtensionSpxSpritePlayAnim(const GDExtensionSpxSpritePlayAnim fn, GdObj obj, GdString p_name, GdFloat p_custom_scale, GdBool p_from_end) {
	fn(obj, p_name, p_custom_scale, p_from_end);
}
void cgo_callfn_GDExtensionSpxSpritePlayBackwardsAnim(const GDExtensionSpxSpritePlayBackwardsAnim fn, GdObj obj, GdString p_name) {
	fn(obj, p_name);
}
void cgo_callfn_GDExtensionSpxSpritePauseAnim(const GDExtensionSpxSpritePauseAnim fn, GdObj obj) {
	fn(obj);
}
void cgo_callfn_GDExtensionSpxSpriteStopAnim(const GDExtensionSpxSpriteStopAnim fn, GdObj obj) {
	fn(obj);
}
void cgo_callfn_GDExtensionSpxSpriteIsPlayingAnim(const GDExtensionSpxSpriteIsPlayingAnim fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetAnim(const GDExtensionSpxSpriteSetAnim fn, GdObj obj, GdString p_name) {
	fn(obj, p_name);
}
void cgo_callfn_GDExtensionSpxSpriteGetAnim(const GDExtensionSpxSpriteGetAnim fn, GdObj obj, GdString* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetAnimFrame(const GDExtensionSpxSpriteSetAnimFrame fn, GdObj obj, GdInt p_frame) {
	fn(obj, p_frame);
}
void cgo_callfn_GDExtensionSpxSpriteGetAnimFrame(const GDExtensionSpxSpriteGetAnimFrame fn, GdObj obj, GdInt* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetAnimSpeedScale(const GDExtensionSpxSpriteSetAnimSpeedScale fn, GdObj obj, GdFloat p_speed_scale) {
	fn(obj, p_speed_scale);
}
void cgo_callfn_GDExtensionSpxSpriteGetAnimSpeedScale(const GDExtensionSpxSpriteGetAnimSpeedScale fn, GdObj obj, GdFloat* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteGetAnimPlayingSpeed(const GDExtensionSpxSpriteGetAnimPlayingSpeed fn, GdObj obj, GdFloat* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetAnimCentered(const GDExtensionSpxSpriteSetAnimCentered fn, GdObj obj, GdBool p_center) {
	fn(obj, p_center);
}
void cgo_callfn_GDExtensionSpxSpriteIsAnimCentered(const GDExtensionSpxSpriteIsAnimCentered fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetAnimOffset(const GDExtensionSpxSpriteSetAnimOffset fn, GdObj obj, GdVec2 p_offset) {
	fn(obj, p_offset);
}
void cgo_callfn_GDExtensionSpxSpriteGetAnimOffset(const GDExtensionSpxSpriteGetAnimOffset fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetAnimFlipH(const GDExtensionSpxSpriteSetAnimFlipH fn, GdObj obj, GdBool p_flip) {
	fn(obj, p_flip);
}
void cgo_callfn_GDExtensionSpxSpriteIsAnimFlippedH(const GDExtensionSpxSpriteIsAnimFlippedH fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetAnimFlipV(const GDExtensionSpxSpriteSetAnimFlipV fn, GdObj obj, GdBool p_flip) {
	fn(obj, p_flip);
}
void cgo_callfn_GDExtensionSpxSpriteIsAnimFlippedV(const GDExtensionSpxSpriteIsAnimFlippedV fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetVelocity(const GDExtensionSpxSpriteSetVelocity fn, GdObj obj, GdVec2 velocity) {
	fn(obj, velocity);
}
void cgo_callfn_GDExtensionSpxSpriteGetVelocity(const GDExtensionSpxSpriteGetVelocity fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteIsOnFloor(const GDExtensionSpxSpriteIsOnFloor fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteIsOnFloorOnly(const GDExtensionSpxSpriteIsOnFloorOnly fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteIsOnWall(const GDExtensionSpxSpriteIsOnWall fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteIsOnWallOnly(const GDExtensionSpxSpriteIsOnWallOnly fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteIsOnCeiling(const GDExtensionSpxSpriteIsOnCeiling fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteIsOnCeilingOnly(const GDExtensionSpxSpriteIsOnCeilingOnly fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteGetLastMotion(const GDExtensionSpxSpriteGetLastMotion fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteGetPositionDelta(const GDExtensionSpxSpriteGetPositionDelta fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteGetFloorNormal(const GDExtensionSpxSpriteGetFloorNormal fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteGetWallNormal(const GDExtensionSpxSpriteGetWallNormal fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteGetRealVelocity(const GDExtensionSpxSpriteGetRealVelocity fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteMoveAndSlide(const GDExtensionSpxSpriteMoveAndSlide fn, GdObj obj) {
	fn(obj);
}
void cgo_callfn_GDExtensionSpxSpriteSetGravity(const GDExtensionSpxSpriteSetGravity fn, GdObj obj, GdFloat gravity) {
	fn(obj, gravity);
}
void cgo_callfn_GDExtensionSpxSpriteGetGravity(const GDExtensionSpxSpriteGetGravity fn, GdObj obj, GdFloat* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetMass(const GDExtensionSpxSpriteSetMass fn, GdObj obj, GdFloat mass) {
	fn(obj, mass);
}
void cgo_callfn_GDExtensionSpxSpriteGetMass(const GDExtensionSpxSpriteGetMass fn, GdObj obj, GdFloat* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteAddForce(const GDExtensionSpxSpriteAddForce fn, GdObj obj, GdVec2 force) {
	fn(obj, force);
}
void cgo_callfn_GDExtensionSpxSpriteAddImpulse(const GDExtensionSpxSpriteAddImpulse fn, GdObj obj, GdVec2 impulse) {
	fn(obj, impulse);
}
void cgo_callfn_GDExtensionSpxSpriteSetCollisionLayer(const GDExtensionSpxSpriteSetCollisionLayer fn, GdObj obj, GdInt layer) {
	fn(obj, layer);
}
void cgo_callfn_GDExtensionSpxSpriteGetCollisionLayer(const GDExtensionSpxSpriteGetCollisionLayer fn, GdObj obj, GdInt* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetCollisionMask(const GDExtensionSpxSpriteSetCollisionMask fn, GdObj obj, GdInt mask) {
	fn(obj, mask);
}
void cgo_callfn_GDExtensionSpxSpriteGetCollisionMask(const GDExtensionSpxSpriteGetCollisionMask fn, GdObj obj, GdInt* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetTriggerLayer(const GDExtensionSpxSpriteSetTriggerLayer fn, GdObj obj, GdInt layer) {
	fn(obj, layer);
}
void cgo_callfn_GDExtensionSpxSpriteGetTriggerLayer(const GDExtensionSpxSpriteGetTriggerLayer fn, GdObj obj, GdInt* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetTriggerMask(const GDExtensionSpxSpriteSetTriggerMask fn, GdObj obj, GdInt mask) {
	fn(obj, mask);
}
void cgo_callfn_GDExtensionSpxSpriteGetTriggerMask(const GDExtensionSpxSpriteGetTriggerMask fn, GdObj obj, GdInt* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetColliderRect(const GDExtensionSpxSpriteSetColliderRect fn, GdObj obj, GdVec2 center, GdVec2 size) {
	fn(obj, center, size);
}
void cgo_callfn_GDExtensionSpxSpriteSetColliderCircle(const GDExtensionSpxSpriteSetColliderCircle fn, GdObj obj, GdVec2 center, GdFloat radius) {
	fn(obj, center, radius);
}
void cgo_callfn_GDExtensionSpxSpriteSetColliderCapsule(const GDExtensionSpxSpriteSetColliderCapsule fn, GdObj obj, GdVec2 center, GdVec2 size) {
	fn(obj, center, size);
}
void cgo_callfn_GDExtensionSpxSpriteSetCollisionEnabled(const GDExtensionSpxSpriteSetCollisionEnabled fn, GdObj obj, GdBool enabled) {
	fn(obj, enabled);
}
void cgo_callfn_GDExtensionSpxSpriteIsCollisionEnabled(const GDExtensionSpxSpriteIsCollisionEnabled fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxSpriteSetTriggerRect(const GDExtensionSpxSpriteSetTriggerRect fn, GdObj obj, GdVec2 center, GdVec2 size) {
	fn(obj, center, size);
}
void cgo_callfn_GDExtensionSpxSpriteSetTriggerCircle(const GDExtensionSpxSpriteSetTriggerCircle fn, GdObj obj, GdVec2 center, GdFloat radius) {
	fn(obj, center, radius);
}
void cgo_callfn_GDExtensionSpxSpriteSetTriggerCapsule(const GDExtensionSpxSpriteSetTriggerCapsule fn, GdObj obj, GdVec2 center, GdVec2 size) {
	fn(obj, center, size);
}
void cgo_callfn_GDExtensionSpxSpriteSetTriggerEnabled(const GDExtensionSpxSpriteSetTriggerEnabled fn, GdObj obj, GdBool trigger) {
	fn(obj, trigger);
}
void cgo_callfn_GDExtensionSpxSpriteIsTriggerEnabled(const GDExtensionSpxSpriteIsTriggerEnabled fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiBindNode(const GDExtensionSpxUiBindNode fn, GdObj obj, GdString rel_path, GdObj* ret_val) {
	fn(obj, rel_path,ret_val);
}
void cgo_callfn_GDExtensionSpxUiCreateNode(const GDExtensionSpxUiCreateNode fn, GdString path, GdObj* ret_val) {
	fn(path,ret_val);
}
void cgo_callfn_GDExtensionSpxUiCreateButton(const GDExtensionSpxUiCreateButton fn, GdString path, GdString text, GdObj* ret_val) {
	fn(path, text,ret_val);
}
void cgo_callfn_GDExtensionSpxUiCreateLabel(const GDExtensionSpxUiCreateLabel fn, GdString path, GdString text, GdObj* ret_val) {
	fn(path, text,ret_val);
}
void cgo_callfn_GDExtensionSpxUiCreateImage(const GDExtensionSpxUiCreateImage fn, GdString path, GdObj* ret_val) {
	fn(path,ret_val);
}
void cgo_callfn_GDExtensionSpxUiCreateToggle(const GDExtensionSpxUiCreateToggle fn, GdString path, GdBool value, GdObj* ret_val) {
	fn(path, value,ret_val);
}
void cgo_callfn_GDExtensionSpxUiCreateSlider(const GDExtensionSpxUiCreateSlider fn, GdString path, GdFloat value, GdObj* ret_val) {
	fn(path, value,ret_val);
}
void cgo_callfn_GDExtensionSpxUiCreateInput(const GDExtensionSpxUiCreateInput fn, GdString path, GdString text, GdObj* ret_val) {
	fn(path, text,ret_val);
}
void cgo_callfn_GDExtensionSpxUiDestroyNode(const GDExtensionSpxUiDestroyNode fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiGetType(const GDExtensionSpxUiGetType fn, GdObj obj, GdInt* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetText(const GDExtensionSpxUiSetText fn, GdObj obj, GdString text) {
	fn(obj, text);
}
void cgo_callfn_GDExtensionSpxUiGetText(const GDExtensionSpxUiGetText fn, GdObj obj, GdString* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetTexture(const GDExtensionSpxUiSetTexture fn, GdObj obj, GdString path) {
	fn(obj, path);
}
void cgo_callfn_GDExtensionSpxUiGetTexture(const GDExtensionSpxUiGetTexture fn, GdObj obj, GdString* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetColor(const GDExtensionSpxUiSetColor fn, GdObj obj, GdColor color) {
	fn(obj, color);
}
void cgo_callfn_GDExtensionSpxUiGetColor(const GDExtensionSpxUiGetColor fn, GdObj obj, GdColor* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetFontSize(const GDExtensionSpxUiSetFontSize fn, GdObj obj, GdInt size) {
	fn(obj, size);
}
void cgo_callfn_GDExtensionSpxUiGetFontSize(const GDExtensionSpxUiGetFontSize fn, GdObj obj, GdInt* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetVisible(const GDExtensionSpxUiSetVisible fn, GdObj obj, GdBool visible) {
	fn(obj, visible);
}
void cgo_callfn_GDExtensionSpxUiGetVisible(const GDExtensionSpxUiGetVisible fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetInteractable(const GDExtensionSpxUiSetInteractable fn, GdObj obj, GdBool interactable) {
	fn(obj, interactable);
}
void cgo_callfn_GDExtensionSpxUiGetInteractable(const GDExtensionSpxUiGetInteractable fn, GdObj obj, GdBool* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetRect(const GDExtensionSpxUiSetRect fn, GdObj obj, GdRect2 rect) {
	fn(obj, rect);
}
void cgo_callfn_GDExtensionSpxUiGetRect(const GDExtensionSpxUiGetRect fn, GdObj obj, GdRect2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiGetLayoutDirection(const GDExtensionSpxUiGetLayoutDirection fn, GdObj obj, GdInt* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetLayoutDirection(const GDExtensionSpxUiSetLayoutDirection fn, GdObj obj, GdInt value) {
	fn(obj, value);
}
void cgo_callfn_GDExtensionSpxUiGetLayoutMode(const GDExtensionSpxUiGetLayoutMode fn, GdObj obj, GdInt* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetLayoutMode(const GDExtensionSpxUiSetLayoutMode fn, GdObj obj, GdInt value) {
	fn(obj, value);
}
void cgo_callfn_GDExtensionSpxUiGetAnchorsPreset(const GDExtensionSpxUiGetAnchorsPreset fn, GdObj obj, GdInt* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetAnchorsPreset(const GDExtensionSpxUiSetAnchorsPreset fn, GdObj obj, GdInt value) {
	fn(obj, value);
}
void cgo_callfn_GDExtensionSpxUiGetScale(const GDExtensionSpxUiGetScale fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetScale(const GDExtensionSpxUiSetScale fn, GdObj obj, GdVec2 value) {
	fn(obj, value);
}
void cgo_callfn_GDExtensionSpxUiGetPosition(const GDExtensionSpxUiGetPosition fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetPosition(const GDExtensionSpxUiSetPosition fn, GdObj obj, GdVec2 value) {
	fn(obj, value);
}
void cgo_callfn_GDExtensionSpxUiGetSize(const GDExtensionSpxUiGetSize fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetSize(const GDExtensionSpxUiSetSize fn, GdObj obj, GdVec2 value) {
	fn(obj, value);
}
void cgo_callfn_GDExtensionSpxUiGetGlobalPosition(const GDExtensionSpxUiGetGlobalPosition fn, GdObj obj, GdVec2* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetGlobalPosition(const GDExtensionSpxUiSetGlobalPosition fn, GdObj obj, GdVec2 value) {
	fn(obj, value);
}
void cgo_callfn_GDExtensionSpxUiGetRotation(const GDExtensionSpxUiGetRotation fn, GdObj obj, GdFloat* ret_val) {
	fn(obj,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetRotation(const GDExtensionSpxUiSetRotation fn, GdObj obj, GdFloat value) {
	fn(obj, value);
}
void cgo_callfn_GDExtensionSpxUiGetFlip(const GDExtensionSpxUiGetFlip fn, GdObj obj, GdBool horizontal, GdBool* ret_val) {
	fn(obj, horizontal,ret_val);
}
void cgo_callfn_GDExtensionSpxUiSetFlip(const GDExtensionSpxUiSetFlip fn, GdObj obj, GdBool horizontal, GdBool is_flip) {
	fn(obj, horizontal, is_flip);
}
#endif
