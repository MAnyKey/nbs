GO_LIBRARY()

LICENSE(MIT)

SRCS(
    backward_references.go
    backward_references_hq.go
    bit_cost.go
    bit_reader.go
    block_splitter.go
    block_splitter_command.go
    block_splitter_distance.go
    block_splitter_literal.go
    brotli_bit_stream.go
    cluster.go
    cluster_command.go
    cluster_distance.go
    cluster_literal.go
    command.go
    compress_fragment.go
    compress_fragment_two_pass.go
    constants.go
    context.go
    decode.go
    dictionary.go
    dictionary_hash.go
    encode.go
    encoder_dict.go
    entropy_encode.go
    entropy_encode_static.go
    fast_log.go
    find_match_length.go
    h10.go
    h5.go
    h6.go
    hash.go
    hash_composite.go
    hash_forgetful_chain.go
    hash_longest_match_quickly.go
    hash_rolling.go
    histogram.go
    http.go
    huffman.go
    literal_cost.go
    memory.go
    metablock.go
    metablock_command.go
    metablock_distance.go
    metablock_literal.go
    params.go
    platform.go
    prefix.go
    prefix_dec.go
    quality.go
    reader.go
    ringbuffer.go
    state.go
    static_dict.go
    static_dict_lut.go
    symbol_list.go
    transform.go
    utf8_util.go
    util.go
    write_bits.go
    writer.go
)

GO_TEST_SRCS(brotli_test.go)

END()

RECURSE(gotest)
