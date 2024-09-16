cmake_minimum_required(VERSION 2.8.12...3.29)

# Core config
set(TARGET_NAME xxhash)
set(LOADABLE_EXTENSION_NAME ${TARGET_NAME}_loadable_extension)
set(GO_EXTENSION_NAME ${TARGET_NAME})
set(GO_SRC ${CMAKE_CURRENT_SOURCE_DIR}/src/go)
set(BUILD_DIR ${CMAKE_BINARY_DIR}/extension/${TARGET_NAME}/${GO_EXTENSION_NAME}-prefix/src/${GO_EXTENSION_NAME}-build)
set(LIB ${BUILD_DIR}/${GO_EXTENSION_NAME}.a)
set(SOURCES src/extension.cpp)

include_directories(src/include)

# Configure extension: the minimum CAPI version where the extension still
# compiles should be chosen here
set(CAPI_MAJOR_VERSION 0)
set(CAPI_MINOR_VERSION 0)
set(CAPI_PATCH_VERSION 1)

set(CAPI_VERSION v${CAPI_MAJOR_VERSION}.${CAPI_MINOR_VERSION}.${CAPI_PATCH_VERSION})

# Set the version definitions. These prevent compilation if a function is used
# that is added in a version > CAPI_VERSION
add_definitions(-DDUCKDB_EXTENSION_API_VERSION_MAJOR=${CAPI_MAJOR_VERSION})
add_definitions(-DDUCKDB_EXTENSION_API_VERSION_MINOR=${CAPI_MINOR_VERSION})
add_definitions(-DDUCKDB_EXTENSION_API_VERSION_PATCH=${CAPI_PATCH_VERSION})
add_definitions(-DDUCKDB_EXTENSION_NAME=${TARGET_NAME})

include(ExternalProject)
ExternalProject_Add(
  ${GO_EXTENSION_NAME}
  SOURCE_DIR ${GO_SRC}
  CONFIGURE_COMMAND ""
  UPDATE_COMMAND ""
  BUILD_COMMAND ${GO_SRC}/build.sh "${GO_SRC}/" "${CMAKE_SOURCE_DIR}/src/include/"
  BUILD_BYPRODUCTS "${LIB}"
  INSTALL_COMMAND ""
  LOG_BUILD ON)

include_directories(${BUILD_DIR})

build_loadable_extension_capi(${TARGET_NAME} ${CAPI_VERSION} ${SOURCES})
#build_static_extension(${TARGET_NAME} ${SOURCES})

target_link_libraries(
  ${LOADABLE_EXTENSION_NAME}
  debug
  "${LIB}"
  optimized
  "${LIB}"
  duckdb)

add_dependencies(${LOADABLE_EXTENSION_NAME} ${GO_EXTENSION_NAME})