# shellcheck disable=SC2148
# UPSTREAM_PKG_IDENT="chef/elasticsearch-odfe/0.10.1.2"
pkg_name="automate-ha-elasticsearch-copy"
pkg_description="Wrapper package for elasticsearch-odfe"
pkg_origin="chef"
vendor_origin="chef"
pkg_version="6.8.14"
pkg_maintainer="Chef Software Inc. <support@chef.io>"
pkg_license=("Chef-MLSA")
pkg_upstream_url="https://www.chef.io/automate"
pkg_source="https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-${pkg_version}.tar.gz"
pkg_shasum=edcf3e951b91fcfb4a7bc0a0f2a58bce74c741dfd64c84c46d9777d75079d5c0
pkg_deps=(
  core/patchelf
  chef/automate-openjdk
)
pkg_lib_dirs=(lib)
pkg_bin_dirs=(es/bin)

pkg_exports=(
  [http-port]=es_yaml.http.port
  [transport-port]=es_yaml.transport.tcp.port
  [root-ca]=opendistro_ssl.rootCA
  [admin-pem]=opendistro_ssl.admin_cert
  [admin-key]=opendistro_ssl.admin_key
  [admin_username]=opendistro_auth.admin_username
  [admin_password]=opendistro_auth.admin_password
  [dashboard_username]=opendistro_auth.dashboard_username
  [dashboard_password]=opendistro_auth.dashboard_password
)
pkg_exposes=(http-port transport-port)

#do_before() {
#  update_pkg_version
#}

do_download() {
  do_default_download
  download_file "https://artifacts.elastic.co/downloads/elasticsearch-plugins/repository-s3/repository-s3-${pkg_version}.zip" "repository-s3.zip" "3dc05d6c20e683596ddabfcc3f63c9d4e9680da75bff1c904566b5508584a6d6"
}

do_build() {
  return 0
}

do_install() {
  cd "$HAB_CACHE_SRC_PATH/elasticsearch-${pkg_version}"
  install -vDm644 README.textile "${pkg_prefix}/README.textile"
  install -vDm644 LICENSE.txt "${pkg_prefix}/LICENSE.txt"
  install -vDm644 NOTICE.txt "${pkg_prefix}/NOTICE.txt"

  # Elasticsearch is greedy when grabbing config files from /bin/..
  # so we need to put the untemplated config dir out of reach
  mkdir -p "${pkg_prefix}/es"
  cp -a ./* "${pkg_prefix}/es"

  # jvm.options needs to live relative to the binary.
  # mkdir -p "$pkg_prefix/es/config"
  # install -vDm644 config/jvm.options "$pkg_prefix/es/config/jvm.options"

  # Delete unused binaries to save space
  rm "${pkg_prefix}/es/bin/"*.bat "${pkg_prefix}/es/bin/"*.exe

  LD_RUN_PATH=$LD_RUN_PATH:${pkg_prefix}/es/modules/x-pack-ml/platform/linux-x86_64/lib
  export LD_RUN_PATH

  for bin in autoconfig autodetect categorize controller normalize; do
    build_line "patch ${pkg_prefix}/es/modules/x-pack-ml/platform/linux-x86_64/bin/${bin}"
    patchelf --interpreter "$(pkg_path_for glibc)/lib/ld-linux-x86-64.so.2" --set-rpath "${LD_RUN_PATH}" \
      "${pkg_prefix}/es/modules/x-pack-ml/platform/linux-x86_64/bin/${bin}"
  done

  find "${pkg_prefix}/es/modules/x-pack-ml/platform/linux-x86_64/lib" -type f -name "*.so" \
      -exec patchelf --set-rpath "${LD_RUN_PATH}" {} \;

  "${pkg_prefix}/es/bin/elasticsearch-plugin" install -b "file://${HAB_CACHE_SRC_PATH}/repository-s3.zip"
}

do_strip() {
  return 0
}
