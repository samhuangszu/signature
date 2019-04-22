#import "NSArray+JSONSorted.h"
#import "NSDictionary+JSONSorted.h"
@implementation NSArray (JSONSorted) 
- (NSString *)sortedJSONString
{
    NSMutableArray *items = [NSMutableArray array];
    for (id item in self) {
        if ([item isKindOfClass:[NSDictionary class]]) {
            [items addObject:[item sortedJSONString]];
        }
        else if ([item isKindOfClass:[NSArray class]]) {
            [items addObject:[item sortedJSONString]];
        }
        else {
            id value = item;
            if ([item isKindOfClass:[NSString class]]) {
                value = [NSString stringWithFormat:@"\"%@\"", item];
            }
            [items addObject:value];
        }
    }
    NSString *value = [NSString stringWithFormat:@"[%@]", [items componentsJoinedByString:@","]];
    return value;
}
@end