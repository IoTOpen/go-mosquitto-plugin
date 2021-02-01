#include <mosquitto.h>
#include <mosquitto_broker.h>
#include <mosquitto_plugin.h>

int goMosquittoPluginVersion(int, int*);
int goMosquittoPluginInit(mosquitto_plugin_id_t*, struct mosquitto_opt*, int);
int goMosquittoPluginCleanup(struct mosquitto_opt*, int);

int goGenericCallback(int, void*, void*);

int mosquitto_plugin_version(int supported_version_count, const int *supported_versions) {
	return goMosquittoPluginVersion(supported_version_count, (int*)supported_versions);
}

int mosquitto_plugin_init(mosquitto_plugin_id_t* identifier, void  **userdata, struct mosquitto_opt *options, int optcount) {
	return goMosquittoPluginInit(identifier, options, optcount);
}

int mosquitto_plugin_cleanup(void *userdata, struct mosquitto_opt *options, int optcount) {
	return goMosquittoPluginCleanup(options, optcount);
}

void go_mosquitto_log_printf(int level, const char* fmt, const char* string) {
    mosquitto_log_printf(level, fmt, string);
}

int go_mosquitto_generic_callback(int event, void* p1, void* p2) {
    return goGenericCallback(event, p1, p2);
}

bool go_mosquitto_topic_matches_sub(char* topic, char* subscription) {
    bool res = false;
    mosquitto_topic_matches_sub(subscription, topic, &res);
    return res;
}

int mosquitto_callback_register2(mosquitto_plugin_id_t* id, int event, void* cb, void* eventData, void* userdata) {
    return mosquitto_callback_register(id, event, cb, eventData, userdata);
}

int mosquitto_callback_unregister2(mosquitto_plugin_id_t* id, int event, void* cb, void* eventData) {
    return mosquitto_callback_unregister(id, event, cb, eventData);
}